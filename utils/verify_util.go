package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// sortedKeys 实现 sort.Interface，用于按 ASCII 升序排序参数键
type sortedKeys []string

func (s sortedKeys) Len() int           { return len(s) }
func (s sortedKeys) Less(i, j int) bool { return s[i] < s[j] }
func (s sortedKeys) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// SignApiRequest 生成淘宝 API 签名（HMAC-SHA256）
// - apiName: API 接口名称
// - params: 请求参数（不包含 sign）
// - secretKey: 应用密钥
// 返回签名字符串（大写 16 进制）
func SignApiRequest(apiName string, params map[string]string, secretKey string) (string, error) {
	// 收集除 sign 外的所有参数名
	var keys []string
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	// 参数名按 ASCII 升序排序
	sort.Sort(sortedKeys(keys))

	// 拼接参数名和参数值
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(params[k])
	}

	// 将 API 名称放在最前面
	stringToSign := apiName + sb.String()

	// 计算 HMAC-SHA256
	h := hmac.New(sha256.New, []byte(secretKey))
	if _, err := h.Write([]byte(stringToSign)); err != nil {
		return "", fmt.Errorf("生成签名失败: %w", err)
	}

	// 转换为大写十六进制字符串
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
}

// VerifySignature 验证推送消息签名
func VerifySignature(r *http.Request, appKey, appSecret string) (bool, string, []byte, error) {
	// 1. 获取 Authorization 头
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false, "", nil, nil
	}

	// 2. 读取请求 body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, "", nil, err
	}
	// 注意：body 读完了需要 reset，否则业务逻辑里就拿不到了
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// 3. 拼接 Base = appKey + body
	base := appKey + string(bodyBytes)

	// 4. 生成 HMAC-SHA256
	h := hmac.New(sha256.New, []byte(appSecret))
	h.Write([]byte(base))
	signature := hex.EncodeToString(h.Sum(nil))

	// 5. 对比签名
	return hmac.Equal([]byte(signature), []byte(authHeader)), signature, bodyBytes, nil
}

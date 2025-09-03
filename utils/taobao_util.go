package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Mlegbder/taobao-global/consts"
	"github.com/Mlegbder/taobao-global/types"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
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

// Execute 调用淘宝 API
// - params: 请求参数（会自动补充通用参数与签名）
// - base: 包含 API 基础信息（如 AppKey、AppSecret、ApiEndpoint）
// 返回响应字节流
func Execute(params map[string]string, base types.TaobaoBase) ([]byte, error) {
	// 添加通用参数
	params["app_key"] = base.AppKey
	params["sign_method"] = consts.TaoBaoApiSignMethod
	params["timestamp"] = strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	signature, err := SignApiRequest(base.ApiEndpoint, params, base.AppSecret)
	if err != nil {
		return nil, fmt.Errorf("签名生成失败: %w", err)
	}
	params["sign"] = signature

	// 转换为 JSON
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("JSON 序列化失败: %w", err)
	}

	// 创建 HTTP POST 请求
	request, err := http.NewRequest("POST", base.Api+base.ApiEndpoint, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("请求创建失败: %w", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	// 发送请求（超时时间 10s）
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("请求发送失败: %w", err)
	}
	defer response.Body.Close()

	// 读取响应体
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 打印原始响应（调试用）
	log.Printf("原始响应内容: %s\n", string(bodyBytes))

	// 检查 HTTP 状态码
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d, 响应: %s", response.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}

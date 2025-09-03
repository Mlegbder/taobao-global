package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
)

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

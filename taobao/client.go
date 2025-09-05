package taobao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Mlegbder/taobao-global/consts"
	"github.com/Mlegbder/taobao-global/types"
	"github.com/Mlegbder/taobao-global/utils"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Client SDK 主入口
type Client struct {
	Base      types.TaobaoBase
	Token     *TokenService
	Item      *ItemService
	Order     *OrderService
	Logistics *LogisticsService
	Upload    *UploadService
	mu        sync.Mutex // 🔒 控制刷新 token 的并发安全

}

// NewClient 创建一个新客户端
func NewClient(baseApi, appKey, appSecret string, tokenStore types.TokenStore) *Client {
	baseConf := types.TaobaoBase{
		AppKey:    appKey,
		AppSecret: appSecret,
		Api:       baseApi,
		Store:     tokenStore,
	}

	client := &Client{Base: baseConf}
	client.Token = &TokenService{client: client}
	client.Item = &ItemService{client: client}
	client.Order = &OrderService{client: client}
	client.Logistics = &LogisticsService{client: client}
	client.Upload = &UploadService{client: client}

	return client
}

func (c *Client) getAccessToken() string {
	tokenStore, _ := c.Base.Store.LoadToken()
	return tokenStore.AccessToken
}

func (c *Client) getRefreshToken() string {
	tokenStore, _ := c.Base.Store.LoadToken()
	return tokenStore.RefreshToken
}

// Execute 执行
func (c *Client) Execute(params map[string]string, base types.TaobaoBase) ([]byte, error) {
	// 先尝试执行一次
	respBytes, err := c.doRequest(params, base)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var resp types.BaseResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	// 判断是否 token 过期
	if resp.Code == consts.IllegalAccessToken {
		log.Println("⚠️ AccessToken 已过期，正在刷新...")

		// 🔒 加锁，避免多个 goroutine 同时刷新
		c.mu.Lock()
		defer c.mu.Unlock()

		// 调用刷新
		var tokenResponse *types.TokenResponse
		tokenResponse, err = c.Token.Refresh(types.RefreshTokenRequest{RefreshToken: c.getRefreshToken()})
		if err != nil {
			return nil, fmt.Errorf("刷新token失败: %w", err)
		}
		err = c.Base.Store.SaveToken(tokenResponse)
		if err != nil {
			return nil, fmt.Errorf("刷新token后储存失败: %w", err)
		}

		// 更新参数里的 token
		params["access_token"] = tokenResponse.AccessToken

		// ✅ 再次调用
		return c.doRequest(params, base)
	}

	if resp.Code != "0" {
		return nil, fmt.Errorf("API 错误: code=%s", resp.Code)
	}
	return respBytes, nil
}

func (c *Client) doRequest(params map[string]string, base types.TaobaoBase) ([]byte, error) {
	// 添加通用参数
	params["app_key"] = base.AppKey
	params["sign_method"] = consts.TaoBaoApiSignMethod
	params["timestamp"] = strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	signature, err := utils.SignApiRequest(base.ApiEndpoint, params, base.AppSecret)
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

	var resp types.BaseResponse
	if err = json.Unmarshal(bodyBytes, &resp); err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

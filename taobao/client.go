package taobao

import (
	"github.com/Mlegbder/taobao-global/types"
)

// Client SDK 主入口
type Client struct {
	Base      types.TaobaoBase
	Token     *TokenService
	Item      *ItemService
	Order     *OrderService
	Logistics *LogisticsService
	Upload    *UploadService
}

// NewClient 创建一个新客户端
func NewClient(baseApi, appKey, appSecret, accessToken string) *Client {
	baseConf := types.TaobaoBase{
		AppKey:      appKey,
		AppSecret:   appSecret,
		Api:         baseApi,
		AccessToken: accessToken,
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
	return c.Base.AccessToken
}

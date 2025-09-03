package taobao

import (
	"taobao-global/consts"
	"taobao-global/types"
)

// Client SDK 主入口
type Client struct {
	Base  types.TaobaoBase
	Token *TokenService
	Item  *ItemService
	Order *OrderService
}

// NewClient 创建一个新客户端
func NewClient(appKey, appSecret string) *Client {
	baseConf := types.TaobaoBase{
		AppKey:    appKey,
		AppSecret: appSecret,
		Api:       consts.BaseApi,
	}

	client := &Client{Base: baseConf}
	client.Token = &TokenService{client: client}
	client.Item = &ItemService{client: client}
	client.Order = &OrderService{client: client}
	return client
}

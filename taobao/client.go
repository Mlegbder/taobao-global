package taobao

import (
	"taobao-global/types"
)

// Client SDK 主入口
type Client struct {
	Base  types.TaobaoBase
	Token *TokenService
	Item  *ItemService
}

// NewClient 创建一个新客户端
func NewClient(appKey, appSecret string) *Client {
	baseConf := types.TaobaoBase{
		AppKey:    appKey,
		AppSecret: appSecret,
		Api:       "https://api.taobao.global/rest",
	}

	client := &Client{Base: baseConf}
	client.Token = &TokenService{client: client}
	client.Item = &ItemService{client: client}

	return client
}

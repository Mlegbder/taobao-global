package taobao

import (
	"encoding/json"
	"taobao-global/consts"
	"taobao-global/types"
	"taobao-global/utils"
)

// TokenService 提供 Token 相关的 API 封装
type TokenService struct {
	client *Client
}

// Create 获取 AccessToken
func (t *TokenService) Create(req types.TokenRequest) (*types.TokenResponse, error) {
	params := map[string]string{
		"code": req.Code,
	}

	baseConf := t.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiGenerateAccessToken

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.TokenResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

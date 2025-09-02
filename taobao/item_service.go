package taobao

import (
	"encoding/json"
	"strconv"
	"taobao-global/consts"
	"taobao-global/types"
	"taobao-global/utils"
)

// ItemService 提供商品相关 API
type ItemService struct {
	client *Client
}

// Search 商品搜索
func (s *ItemService) Search(req types.ItemSearchRequest, accessToken string) (*types.ItemSearchResponse, error) {
	params := map[string]string{
		"access_token": accessToken, // 必须传入 access_token
	}

	if req.Keyword != "" {
		params["keyword"] = req.Keyword
	}
	if len(req.IncludeTags) > 0 {
		params["include_tags"] = joinStrings(req.IncludeTags, ",")
	}
	if req.Sort != "" {
		params["sort"] = req.Sort
	}
	if req.PageNo > 0 {
		params["page_no"] = strconv.Itoa(req.PageNo)
	}
	if req.PageSize > 0 {
		params["page_size"] = strconv.Itoa(req.PageSize)
	}
	if len(req.Filters) > 0 {
		params["filters"] = joinStrings(req.Filters, ",")
	}
	if req.Language != "" {
		params["language"] = req.Language
	}
	if req.ShopID > 0 {
		params["shop_id"] = strconv.Itoa(req.ShopID)
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiProductSearch

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.ItemSearchResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// joinStrings 辅助方法
func joinStrings(arr []string, sep string) string {
	out := ""
	for i, v := range arr {
		if i > 0 {
			out += sep
		}
		out += v
	}
	return out
}

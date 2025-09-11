package taobao

import (
	"encoding/json"
	"github.com/Mlegbder/taobao-global/consts"
	"github.com/Mlegbder/taobao-global/types"
	"strconv"
)

// ItemService 提供商品相关 API
type ItemService struct {
	client *Client
}

// Search 商品搜索
func (s *ItemService) Search(req types.ItemSearchRequest) (*types.ItemSearchResponse, error) {
	params := map[string]string{
		"access_token": s.client.getAccessToken(), // 必须传入 access_token
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

	respBytes, err := s.client.Execute(params, baseConf)
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

// GetDetail 获取商品详情
func (s *ItemService) GetDetail(req types.QueryAllProductRequest) (*types.QueryAllProductResponse, error) {
	params := map[string]string{
		"access_token": s.client.getAccessToken(),
		"item_id":      req.ItemID,
	}
	if req.ItemSourceMarket != "" {
		params["item_source_market"] = req.ItemSourceMarket
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiQueryAllProduct

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.QueryAllProductResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSourceItemDetail 获取商品详情
func (s *ItemService) GetSourceItemDetail(req types.ItemDetailRequest) (*types.ItemDetailResponse, error) {
	params := map[string]string{
		"access_token":  s.client.getAccessToken(),
		"item_resource": req.ItemResource,
		"item_id":       req.ItemID,
	}

	if len(req.IncludeTags) > 0 {
		params["include_tags"] = joinStrings(req.IncludeTags, ",")
	}
	if req.Language != "" {
		params["language"] = req.Language
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiSourceItemDetail

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.ItemDetailResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImgSearch 图片搜索
func (s *ItemService) ImgSearch(req types.ImgSearchRequest) (*types.ImgSearchResponse, error) {
	params := map[string]string{
		"access_token": s.client.getAccessToken(),
	}

	if req.PicURL != "" {
		params["pic_url"] = req.PicURL
	}
	if len(req.IncludeTags) > 0 {
		params["include_tags"] = joinStrings(req.IncludeTags, ",")
	}
	if req.Language != "" {
		params["language"] = req.Language
	}
	if req.ImageID != "" {
		params["image_id"] = req.ImageID
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiImgSearch

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.ImgSearchResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Translate 商品信息翻译
func (s *ItemService) Translate(req types.ProductTranslateRequest) (*types.ProductInfoTranResponse, error) {
	params := map[string]string{
		"access_token": s.client.getAccessToken(),
		"item_id":      req.ItemID,
	}
	if req.Language != "" {
		params["language"] = req.Language
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiProductInfoTran

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.ProductInfoTranResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

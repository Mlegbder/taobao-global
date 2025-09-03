package taobao

import (
	"encoding/json"
	"strconv"
	"taobao-global/consts"
	"taobao-global/types"
	"taobao-global/utils"
)

// LogisticsService 提供物流相关API
type LogisticsService struct {
	client *Client
}

// GetDetail 获取子单物流信息
func (s *LogisticsService) GetDetail(req types.GetLogisticsDetailRequest, accessToken string) (*types.GetLogisticsDetailResponse, error) {
	params := map[string]string{
		"access_token":           accessToken,
		"purchase_order_line_id": strconv.FormatInt(req.PurchaseOrderLineID, 10),
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiLogisticsDetail

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.GetLogisticsDetailResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

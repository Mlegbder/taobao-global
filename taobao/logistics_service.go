package taobao

import (
	"encoding/json"
	"github.com/Mlegbder/taobao-global/consts"
	"github.com/Mlegbder/taobao-global/types"
	"strconv"
)

// LogisticsService 提供物流相关API
type LogisticsService struct {
	client *Client
}

// GetDetail 获取子单物流信息
func (s *LogisticsService) GetDetail(req types.GetLogisticsDetailRequest) (*types.LogisticsDetailResponse, error) {
	params := map[string]string{
		"access_token":           s.client.getAccessToken(),
		"purchase_order_line_id": strconv.FormatInt(req.PurchaseOrderLineID, 10),
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiLogisticsDetail

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.LogisticsDetailResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

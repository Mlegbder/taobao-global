package taobao

import (
	"encoding/json"
	"github.com/Mlegbder/taobao-global/consts"
	"github.com/Mlegbder/taobao-global/types"
	"strconv"
)

// BillService 提供账单相关 API
type BillService struct {
	client *Client
}

// PurchaseBill 采购账单查询
func (s *BillService) PurchaseBill(req types.PurchaseBillRequest) (*types.PurchaseBillResponse, error) {
	params := map[string]string{
		"access_token": s.client.getAccessToken(), // 必须传入 access_token
		"time_type":    req.TimeType,
		"start_time":   strconv.FormatInt(req.StartTime, 10),
		"end_time":     strconv.FormatInt(req.EndTime, 10),
		"page_no":      strconv.Itoa(req.PageNo),
		"page_size":    strconv.Itoa(req.PageSize),
	}

	if req.PurchaseOrderID != "" {
		params["purchase_order_id"] = req.PurchaseOrderID
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiPurchaseBill

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.PurchaseBillResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefundBill 退款账单查询
func (s *BillService) RefundBill(req types.RefundBillRequest) (*types.RefundBillResponse, error) {
	params := map[string]string{
		"access_token": s.client.getAccessToken(), // 必须传入 access_token
		"start_time":   req.StartTime,
		"end_time":     req.EndTime,
	}

	if req.PurchaseOrderID != "" {
		params["purchase_order_id"] = req.PurchaseOrderID
	}
	if req.TaobaoOrderID != "" {
		params["taobao_order_id"] = req.TaobaoOrderID
	}
	if req.PageNo != "" {
		params["page_no"] = req.PageNo
	}
	if req.PageSize != "" {
		params["page_size"] = req.PageSize
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiRefundBill

	respBytes, err := s.client.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.RefundBillResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

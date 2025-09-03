package taobao

import (
	"encoding/json"
	"github.com/Mlegbder/taobao-global/consts"
	"github.com/Mlegbder/taobao-global/types"
	"github.com/Mlegbder/taobao-global/utils"
	"strconv"
)

// OrderService 提供订单相关API
type OrderService struct {
	client *Client
}

// Render 订单预览
func (s *OrderService) Render(req types.PurchaseOrderRenderRequest, accessToken string) (*types.PurchaseOrderRenderResponse, error) {
	params := map[string]string{
		"access_token":             accessToken,
		"need_supplychain_service": strconv.FormatBool(req.NeedSupplyChainService),
		"render_item_List":         req.RenderItemList,
	}

	// warehouse_address 可选
	if req.WarehouseAddress != nil {
		b, _ := json.Marshal(req.WarehouseAddress)
		params["warehouse_address"] = string(b)
	}

	// receiver_address 必填
	b, _ := json.Marshal(req.ReceiverAddress)
	params["receiver_address"] = string(b)

	if req.TaxID != "" {
		params["tax_id"] = req.TaxID
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiPurchaseOrderRender

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.PurchaseOrderRenderResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create 创建采购订单
func (s *OrderService) Create(req types.CreatePurchaseOrderRequest, accessToken string) (*types.CreatePurchaseOrderResponse, error) {
	params := map[string]string{
		"access_token":      accessToken,
		"outer_purchase_id": req.OuterPurchaseID,
		"purchase_amount":   strconv.FormatInt(req.PurchaseAmount, 10),
		"order_line_list":   req.OrderLineList,
	}

	if req.SellerOrderNumber != "" {
		params["seller_order_number"] = req.SellerOrderNumber
	}
	if req.OrderSource != "" {
		params["order_source"] = req.OrderSource
	}
	if req.ChannelOrderType != "" {
		params["channel_order_type"] = req.ChannelOrderType
	}
	if req.OrderRemark != "" {
		params["order_remark"] = req.OrderRemark
	}
	if req.SupportPartialSuccess {
		params["support_partial_success"] = "true"
	}
	if !req.NeedSysRetry {
		params["need_sys_retry"] = "false"
	}

	// receiver (必填)
	receiverBytes, _ := json.Marshal(req.Receiver)
	params["receiver"] = string(receiverBytes)

	// warehouse_address_info (可选)
	if req.WarehouseAddressInfo != nil {
		waBytes, _ := json.Marshal(req.WarehouseAddressInfo)
		params["warehouse_address_info"] = string(waBytes)
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiCreatePurchaseOrder

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.CreatePurchaseOrderResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AsynCancel 异步取消采购订单
func (s *OrderService) AsynCancel(req types.AsynCancelPurchaseOrderRequest, accessToken string) (*types.AsynCancelPurchaseOrderResponse, error) {
	params := map[string]string{
		"access_token":  accessToken,
		"purchase_id":   req.PurchaseID,
		"cancel_reason": req.CancelReason,
	}

	if len(req.SubPurchaseOrderIDs) > 0 {
		b, _ := json.Marshal(req.SubPurchaseOrderIDs)
		params["sub_purchase_orderId_list"] = string(b)
	}
	if req.CancelRemark != "" {
		params["cancel_remark"] = req.CancelRemark
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiOrderCancel

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.AsynCancelPurchaseOrderResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchPay 批量支付采购订单
func (s *OrderService) BatchPay(req types.BatchPayPurchaseOrderRequest, accessToken string) (*types.BatchPayPurchaseOrderResponse, error) {
	// 将数组序列化成 JSON 字符串
	b, _ := json.Marshal(req.PurchaseOrderIDList)

	params := map[string]string{
		"access_token":        accessToken,
		"purchaseOrderIdList": string(b),
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiPurchaseOrderBatchPay

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.BatchPayPurchaseOrderResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Query 采购订单查询
func (s *OrderService) Query(req types.QueryPurchaseOrdersRequest, accessToken string) (*types.QueryPurchaseOrdersResponse, error) {
	params := map[string]string{
		"access_token": accessToken,
	}

	if req.Status != "" {
		params["status"] = req.Status
	}
	if req.SortType != "" {
		params["sort_type"] = req.SortType
	}
	if req.PageNo > 0 {
		params["page_no"] = strconv.Itoa(req.PageNo)
	}
	if req.PageSize > 0 {
		params["page_size"] = strconv.Itoa(req.PageSize)
	}
	if req.ModifyTimeStart > 0 {
		params["modify_time_start"] = strconv.FormatInt(req.ModifyTimeStart, 10)
	}
	if req.ModifyTimeEnd > 0 {
		params["modify_time_end"] = strconv.FormatInt(req.ModifyTimeEnd, 10)
	}
	if req.OuterPurchaseID != "" {
		params["outer_purchase_id"] = req.OuterPurchaseID
	}
	if len(req.PurchaseIDS) > 0 {
		b, _ := json.Marshal(req.PurchaseIDS)
		params["purchase_ids"] = string(b)
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiOrdersQuery

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.QueryPurchaseOrdersResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

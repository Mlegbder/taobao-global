package taobao

import (
	"encoding/json"
	"strconv"
	"taobao-global/consts"
	"taobao-global/types"
	"taobao-global/utils"
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

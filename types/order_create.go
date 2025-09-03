package types

// CreatePurchaseOrderRequest 采购订单创建请求
type CreatePurchaseOrderRequest struct {
	OuterPurchaseID        string        `json:"outer_purchase_id"`                   // 必填: ISV采购id (幂等)
	PurchaseAmount         int64         `json:"purchase_amount"`                     // 必填: 计划购买金额 (单位分)
	SellerOrderNumber      string        `json:"seller_order_number,omitempty"`       // 可选: 买方电商平台订单号
	OrderSource            string        `json:"order_source,omitempty"`              // 可选: 买方电商平台
	OrderLineList          string        `json:"order_line_list"`                     // 必填: 下单商品信息列表 (itemId传mpId, skuId传mpSkuId)
	Receiver               OrderAddress  `json:"receiver"`                            // 必填: 收货地址
	WarehouseAddressInfo   *OrderAddress `json:"warehouse_address_info,omitempty"`    // 可选: 仓库地址 (非跨境必填)
	ChannelOrderType       string        `json:"channel_order_type,omitempty"`        // 可选: PANAMA_DG / PANAMA
	SupportPartialSuccess  bool          `json:"support_partial_success,omitempty"`   // 可选: 是否允许部分成功
	OrderRemark            string        `json:"order_remark,omitempty"`              // 可选: 订单备注
	NeedSupplyChainService bool          `json:"need_supply_chain_service,omitempty"` // 暂不提供
	NeedSysRetry           bool          `json:"need_sys_retry,omitempty"`            // 可选: 是否系统重试
}

// OrderAddress 地址
type OrderAddress struct {
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
	City        string `json:"city"`
	District    string `json:"district,omitempty"`
	Name        string `json:"name"`
	State       string `json:"state"`
	MobilePhone string `json:"mobile_phone,omitempty"`
	TaxID       string `json:"taxId,omitempty"`
}

// CreatePurchaseOrderResponse 采购订单创建响应
type CreatePurchaseOrderResponse struct {
	Success   bool               `json:"success"`
	ErrorCode string             `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
	Data      *PurchaseOrderData `json:"data"`
}

// PurchaseOrderData 响应业务数据
type PurchaseOrderData struct {
	PaymentURL      string      `json:"payment_url"`
	OuterPurchaseID string      `json:"outer_purchase_id"`
	Success         bool        `json:"success"`
	ErrorMessage    string      `json:"error_message"`
	ErrorCode       string      `json:"error_code"`
	OrderList       []OrderInfo `json:"order_list"`
}

// OrderInfo 主订单
type OrderInfo struct {
	PurchaseID       int64       `json:"purchase_id"`
	SupplierNick     string      `json:"supplier_nick"`
	EstimateAmount   int64       `json:"estimate_amount"`
	EstimateCurrency string      `json:"estimate_currency"`
	SourceMarket     string      `json:"source_market"`
	OrderLineList    []OrderLine `json:"order_line_list"`
}

// OrderLine 子订单
type OrderLine struct {
	ItemID             string `json:"item_id"`
	OrderLineNo        string `json:"order_line_no"`
	Quantity           int    `json:"quantity"`
	EstimateAmount     int64  `json:"estimate_amount"`
	EstimateCurrency   string `json:"estimate_currency"`
	SubPurchaseOrderID int64  `json:"sub_purchase_order_id"`
	SkuID              int64  `json:"sku_id"`
	SupplierNick       string `json:"supplier_nick"`
}

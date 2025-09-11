package types

// RefundBillRequest 退款账单查询请求参数
type RefundBillRequest struct {
	StartTime       string `json:"start_time"`                  // 开始时刻 【必填】
	EndTime         string `json:"end_time"`                    // 结束时刻 【必填】
	PurchaseOrderID string `json:"purchase_order_id,omitempty"` // 采购单ID，可选
	TaobaoOrderID   string `json:"taobao_order_id,omitempty"`   // 淘宝订单ID，可选
	PageNo          string `json:"page_no,omitempty"`           // 页码，可选
	PageSize        string `json:"page_size,omitempty"`         // 页面大小，可选
}

// RefundBillResponse API: 查询退款账单
type RefundBillResponse struct {
	Result RefundBillResult `json:"result"` // 结果集
}

// RefundBillResult 结果集
type RefundBillResult struct {
	Success       bool           `json:"success"`         // 成功标
	FailItems     string         `json:"fail_items"`      // 失败条目
	ErrorCode     string         `json:"error_code"`      // 错误码
	ErrorMsg      string         `json:"error_msg"`       // 错误信息
	InnerErrorMsg string         `json:"inner_error_msg"` // 内部错误详情
	Data          RefundBillData `json:"data"`            // 业务结果
}

// RefundBillData 业务结果（分页 + 数据列表）
type RefundBillData struct {
	PageNo     int              `json:"page_no"`     // 当前页码
	TotalPage  int              `json:"total_page"`  // 总页数
	PageSize   int              `json:"page_size"`   // 每页大小
	TotalCount int              `json:"total_count"` // 总条数
	Data       []RefundBillItem `json:"data"`        // 退款账单列表
}

// RefundBillItem 单条退款账单
type RefundBillItem struct {
	PurchaseMordID        string `json:"purchase_mord_id"`         // 采购主单ID
	PurchaseOrderID       string `json:"purchase_order_id"`        // 采购子单ID
	PurchaseItemTitleOut  string `json:"purchase_item_title_out"`  // 客户传入商品名称
	RefundChannelTradeNo  string `json:"refund_channel_trade_no"`  // 退款渠道交易号
	OuterPurchaseTime     string `json:"outer_purchase_time"`      // 下单时间（时间戳）
	PurchaseRefundAmount  string `json:"purchase_refund_amount"`   // 订单退款金额（CNY，单位：分）
	PayTime               string `json:"pay_time"`                 // 支付时间（时间戳）
	PayCurrency           string `json:"pay_currency"`             // 支付币种
	PoOrderID             string `json:"po_order_id"`              // 淘宝订单ID
	PurchaseItemTitleReal string `json:"purchase_item_title_real"` // 实际商品名称
	OuterPurchaseID       string `json:"outer_purchase_id"`        // 渠道订单ID
	RefundStartTime       string `json:"refund_start_time"`        // 申请退款时间（时间戳）
	RefundEndTime         string `json:"refund_end_time"`          // 退款完结时间（时间戳）
	RefundCurrency        string `json:"refund_currency"`          // 实际退款币种
	PayChannel            string `json:"pay_channel"`              // 支付方式 (Alipay/Ipay)
	RealRefundAmount      string `json:"real_refund_amount"`       // 实际退款金额（单位：分）
	PayChannelTradeNo     string `json:"pay_channel_trade_no"`     // 支付渠道交易号
	PurchaseItemID        string `json:"purchase_item_id"`         // 采购商品ID
	PoMordID              string `json:"po_mord_id"`               // 货源市场主单ID
}

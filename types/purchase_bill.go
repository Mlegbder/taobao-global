package types

// PurchaseBillRequest 采购账单查询请求参数
type PurchaseBillRequest struct {
	TimeType        string `json:"time_type"`                   // 时间类型 (createtime / paytime) 【必填】
	PurchaseOrderID string `json:"purchase_order_id,omitempty"` // 采购子单ID，可选
	PageNo          int    `json:"page_no"`                     // 页码 【必填】
	PageSize        int    `json:"page_size"`                   // 页面大小 【必填】
	StartTime       int64  `json:"start_time"`                  // 开始时间 (时间戳) 【必填】
	EndTime         int64  `json:"end_time"`                    // 结束时间 (时间戳) 【必填】
}

// PurchaseBillResponse API: 查询采购账单
type PurchaseBillResponse struct {
	Result PurchaseBillResult `json:"result"` // 结果集
}

// PurchaseBillResult 结果集
type PurchaseBillResult struct {
	InnerErrorMsg string           `json:"inner_error_msg"` // 内部错误详情
	Data          PurchaseBillData `json:"data"`            // 业务结果
	Success       bool             `json:"success"`         // 成功标
	ErrorCode     string           `json:"error_code"`      // 错误码
	ErrorMsg      string           `json:"error_msg"`       // 错误信息
}

// PurchaseBillData 业务结果
type PurchaseBillData struct {
	BizCode    string             `json:"biz_code"`    // 成功码
	Data       []PurchaseBillItem `json:"data"`        // 业务结果列表
	PageNo     int                `json:"page_no"`     // 当前页面
	TotalPage  int                `json:"total_page"`  // 总页数
	Success    bool               `json:"success"`     // 成功标
	PageSize   int                `json:"page_size"`   // 每页大小
	TotalCount int                `json:"total_count"` // 总条数
	BizMessage string             `json:"biz_message"` // 业务信息
}

// PurchaseBillItem 单条采购账单
type PurchaseBillItem struct {
	PurchaseDiscountAmount  string `json:"purchase_discount_amount"`  // 订单优惠金额（元）
	PurchaseItemID          string `json:"purchase_item_id"`          // 采购商品ID
	PurchaseCouponAmount    string `json:"purchase_coupon_amount"`    // 订单红包抵扣金额（元）
	PurchaseOrderID         string `json:"purchase_order_id"`         // 采购子单ID
	MainlandPostAmount      string `json:"mainland_post_amount"`      // 订单大陆运费金额（元）
	PaymentServiceCharge    string `json:"payment_service_charge"`    // 订单支付手续费（元）
	PurchaseRefundAmount    string `json:"purchase_refund_amount"`    // 订单退款金额（元）
	PayCurrency             string `json:"pay_currency"`              // 采购单原币种
	PurchaseItemTitleReal   string `json:"purchase_item_title_real"`  // 实际商品名称
	PurchaseAmountEstimated string `json:"purchase_amount_estimated"` // 订单预估采购金额（元）
	PurchaseItemQuantity    string `json:"purchase_item_quantity"`    // 采购商品数量
	OuterPurchaseID         string `json:"outer_purchase_id"`         // 渠道订单ID
	RefundEndTime           string `json:"refund_end_time"`           // 退款完结时间（时间戳）
	PurchaseAmountOriginal  string `json:"purchase_amount_original"`  // 订单原始采购金额（元）
	PayChannelTradeNo       string `json:"pay_channel_trade_no"`      // 支付渠道交易号
	PurchasePayAmount       string `json:"purchase_pay_amount"`       // 订单实付金额（元）
	ExchangeRate            string `json:"exchange_rate"`             // 币种转换汇率
	PoMordID                string `json:"po_mord_id"`                // 货源市场主单ID
	PurchaseMordID          string `json:"purchase_mord_id"`          // 采购主单ID
	PurchaseCurrency        string `json:"purchase_currency"`         // 采购单原币种
	PurchaseAdjustAmount    string `json:"purchase_adjust_amount"`    // 订单卖家调整金额（元）
	PayAmountTotal          string `json:"pay_amount_total"`          // 总支付金额（元）
	PurchaseItemTitleOut    string `json:"purchase_item_title_out"`   // 客户传入商品名称
	OuterPurchaseTime       string `json:"outer_purchase_time"`       // 渠道订单创建时间（时间戳）
	PayTime                 int64  `json:"pay_time"`                  // 支付时间（时间戳）
	PoOrderID               string `json:"po_order_id"`               // 货源市场子单ID
	PurchaseTotalPrice      string `json:"purchase_total_price"`      // 订单应付金额（元）
	RefundStartTime         int64  `json:"refund_start_time"`         // 退款申请时间（时间戳）
	PayChannel              string `json:"pay_channel"`               // 支付渠道
}

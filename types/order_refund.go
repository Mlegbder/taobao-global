package types

// QueryRefundOrderRequest 请求参数
type QueryRefundOrderRequest struct {
	RefundID int64 `json:"refundId"` // refund Id / 逆向单id（必填）
}

// QueryRefundOrderResponse API: QueryRefundOrder
type QueryRefundOrderResponse struct {
	Result QueryRefundOrderResult `json:"result"` // 结果集
}

// QueryRefundOrderResult 结果集
type QueryRefundOrderResult struct {
	Data          QueryRefundOrderData `json:"data"`            // 业务结果
	InnerErrorMsg string               `json:"inner_error_msg"` // 内部错误详情
	Success       bool                 `json:"success"`         // 成功标
	FailItems     string               `json:"fail_items"`      // 错误条目
	ErrorCode     string               `json:"error_code"`      // 错误码
	ErrorMsg      string               `json:"error_msg"`       // 错误信息
}

// QueryRefundOrderData 业务结果
type QueryRefundOrderData struct {
	DistributorID        int64                     `json:"distributor_id"`         // 分销商id
	RefundOrderAggregate QueryRefundOrderAggregate `json:"refund_order_aggregate"` // 逆向聚合信息
	Currency             string                    `json:"currency"`               // 币种
}

// QueryRefundOrderAggregate 逆向聚合信息
type QueryRefundOrderAggregate struct {
	Logistics         RefundLogistics         `json:"logistics"`           // 物流信息
	RefundOrder       RefundOrder             `json:"refund_order"`        // 逆向单基础信息
	PurchaseOrderLine RefundPurchaseOrderLine `json:"purchase_order_line"` // 采购单信息
}

// RefundLogistics 物流信息
type RefundLogistics struct {
	ReceiverAddress           ReceiverAddress  `json:"receiver_address"`             // 收货地址
	ReturnGoodsDesc           string           `json:"return_goods_desc"`            // 退货描述信息
	BuyerPhone                string           `json:"buyer_phone"`                  // 买家手机号
	LogisticsCompany          LogisticsCompany `json:"logistics_company"`            // 物流公司
	SellerAgreeReturnDescribe string           `json:"seller_agree_return_describe"` // 卖家同意退款时的说明
	LogisticsNo               string           `json:"logistics_no"`                 // 物流单号
}

// ReceiverAddress 收货地址
type ReceiverAddress struct {
	Receiver      string `json:"receiver"`       // 收货人
	Phone         string `json:"phone"`          // 收货手机号
	DetailAddress string `json:"detail_address"` // 收货地址
}

// LogisticsCompany 物流公司
type LogisticsCompany struct {
	LogisticsCompanyName string `json:"logistics_company_name"` // 物流公司编号
	LogisticsCompanyCode string `json:"logistics_company_code"` // 物流公司名称
}

// RefundOrder 逆向单基础信息
type RefundOrder struct {
	RefundFee                  int64  `json:"refund_fee"`                    // 逆向金额
	PayOrderID                 string `json:"pay_order_id"`                  // 逆向打款的支付单号
	TimeoutType                string `json:"timeout_type"`                  // 超时类型
	SellerRefuseReason         string `json:"seller_refuse_reason"`          // 卖家拒绝退款申请的原因
	RefundStatus               int    `json:"refund_status"`                 // 逆向状态
	CanEditFee                 string `json:"can_edit_fee"`                  // 能否编辑退款金额
	RemainingTime              int64  `json:"remaining_time"`                // 逆向剩余时间
	RefundType                 int    `json:"refund_type"`                   // 退款类型
	FeeTips                    string `json:"fee_tips"`                      // 退款金额提示
	RefundDesc                 string `json:"refund_desc"`                   // 退款原因模板描述
	GoodsStatus                int    `json:"goods_status"`                  // 货物状态
	ReasonID                   int64  `json:"reason_id"`                     // 逆向原因模板id
	TimeoutDate                int64  `json:"timeout_date"`                  // 逆向超时时间
	SellerRefuseReasonDescribe string `json:"seller_refuse_reason_describe"` // 卖家拒绝申请退款的说明
	ReasonDesc                 string `json:"reason_desc"`                   // 退款原因模板描述
	RefundID                   int64  `json:"refund_id"`                     // 逆向单id
}

// RefundPurchaseOrderLine 采购单信息
type RefundPurchaseOrderLine struct {
	SalesOrderID            string              `json:"sales_order_id"`            // 销售主单id
	SupplierID              int64               `json:"supplier_id"`               // 供应商id
	Quantity                int                 `json:"quantity"`                  // 商品数量
	TotalPrice              int64               `json:"total_price"`               // 订单总金额
	ItemTitle               string              `json:"item_title"`                // 商品标题
	PostFee                 int64               `json:"post_fee"`                  // 运费
	ProductFee              int64               `json:"product_fee"`               // 商品总价
	SalesMarket             string              `json:"sales_market"`              // 采购市场
	PurchaseOrderLineID     int64               `json:"purchase_order_line_id"`    // 采购子单id
	CreateTime              int64               `json:"create_time"`               // 采购子单创建时间
	Price                   string              `json:"price"`                     // 商品单价
	PurchaseOrderID         int64               `json:"purchase_order_id"`         // 采购单主单id
	SupplierNick            string              `json:"supplier_nick"`             // 供应商nick
	Currency                string              `json:"currency"`                  // 币种
	Properties              RefundProperties    `json:"properties"`                // 商品属性
	ItemImageURL            string              `json:"item_image_url"`            // 商品图片url
	OrderServiceExpressions []map[string]string `json:"order_service_expressions"` // 服务信息
	HongbaoDeductionAmount  int64               `json:"hongbao_deduction_amount"`  // 退回的红包金额
}

// RefundProperties 商品属性
type RefundProperties struct {
	Options string `json:"options"` // 属性内容
}

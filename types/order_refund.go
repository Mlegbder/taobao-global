package types

// QueryRefundOrderRequest 查询退款单请求
type QueryRefundOrderRequest struct {
	RefundID int64 `json:"refundId"` // 必填: 退款单ID (逆向单id)
}

// QueryRefundOrderResponse 查询退款单响应
type QueryRefundOrderResponse struct {
	Result        *RefundResult    `json:"result"`          // 结果集
	Data          *RefundOrderData `json:"data"`            // 业务结果
	InnerErrorMsg string           `json:"inner_error_msg"` // 内部错误详情
	Success       bool             `json:"success"`         // 是否成功
	FailItems     string           `json:"fail_items"`      // 错误项
	ErrorCode     string           `json:"error_code"`      // 错误码
	ErrorMsg      string           `json:"error_msg"`       // 错误信息
}

// RefundResult 通用结果
type RefundResult struct {
	Success   bool   `json:"success"`    // 成功标志
	ErrorCode string `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误信息
}

// RefundOrderData 退款单详情
type RefundOrderData struct {
	DistributorID          int64                    `json:"distributor_id"`           // 分销商id
	RefundOrderAggregate   *RefundOrderAggregate    `json:"refund_order_aggregate"`   // 逆向聚合信息
	RefundOrder            *RefundOrderInfo         `json:"refund_order"`             // 退款单基础信息
	PurchaseOrderLine      *RefundPurchaseOrderLine `json:"purchase_order_line"`      // 采购单信息
	HongbaoDeductionAmount int64                    `json:"hongbao_deduction_amount"` // 退回的红包金额 (单位分)
	Currency               string                   `json:"currency"`                 // 币种
}

// RefundOrderAggregate 逆向聚合信息
type RefundOrderAggregate struct {
	Logistics                 *RefundLogistics `json:"logistics"`                    // 物流信息
	ReceiverAddress           *RefundAddress   `json:"receiver_address"`             // 收货地址
	ReturnGoodsDesc           string           `json:"return_goods_desc"`            // 退货描述信息
	BuyerPhone                string           `json:"buyer_phone"`                  // 买家手机号
	SellerAgreeReturnDescribe string           `json:"seller_agree_return_describe"` // 卖家同意退款说明
	LogisticsNo               string           `json:"logistics_no"`                 // 物流单号
}

// RefundLogistics 物流公司信息
type RefundLogistics struct {
	LogisticsCompanyName string `json:"logistics_company_name"` // 物流公司名称
	LogisticsCompanyCode string `json:"logistics_company_code"` // 物流公司编号
}

// RefundAddress 收货地址
type RefundAddress struct {
	Receiver      string `json:"receiver"`       // 收货人
	Phone         string `json:"phone"`          // 收货手机号
	DetailAddress string `json:"detail_address"` // 收货地址
}

// RefundOrderInfo 基础退款单信息
type RefundOrderInfo struct {
	RefundFee                  int64  `json:"refund_fee"`                    // 退款金额 (单位分)
	PayOrderID                 string `json:"pay_order_id"`                  // 打款的支付单号
	TimeoutType                string `json:"timeout_type"`                  // 超时类型
	SellerRefuseReason         string `json:"seller_refuse_reason"`          // 卖家拒绝退款原因
	RefundStatus               int    `json:"refund_status"`                 // 退款状态 (0:未申请, 10:等待卖家同意, 20:等待买家退货, 30:等待卖家确认收货, 100:成功, -10:卖家拒绝, -20:关闭)
	CanEditFee                 string `json:"can_edit_fee"`                  // 是否可修改退款金额
	RemainingTime              string `json:"remaining_time"`                // 剩余时间
	RefundType                 int    `json:"refund_type"`                   // 退款类型 (1:仅退款, 2:退货退款)
	FeeTips                    string `json:"fee_tips"`                      // 金额提示
	RefundDesc                 string `json:"refund_desc"`                   // 退款原因描述
	GoodsStatus                int    `json:"goods_status"`                  // 货物状态 (1:未发货, 2:已发货, 3:未收到货, 4:已收到货, 5:已寄回, 6:卖家确认收货)
	ReasonID                   int64  `json:"reason_id"`                     // 退款原因模板id
	TimeoutDate                int64  `json:"timeout_date"`                  // 超时时间
	SellerRefuseReasonDescribe string `json:"seller_refuse_reason_describe"` // 卖家拒绝退款说明
	ReasonDesc                 string `json:"reason_desc"`                   // 退款原因模板描述
	RefundID                   int64  `json:"refund_id"`                     // 退款单id
}

// RefundPurchaseOrderLine 采购单信息
type RefundPurchaseOrderLine struct {
	SalesOrderID        string            `json:"sales_order_id"`         // 销售主单id
	SupplierID          int64             `json:"supplier_id"`            // 供应商id
	Quantity            int               `json:"quantity"`               // 商品数量
	TotalPrice          int64             `json:"total_price"`            // 订单总金额 (单位分)
	ItemTitle           string            `json:"item_title"`             // 商品标题
	PostFee             int64             `json:"post_fee"`               // 运费 (单位分)
	ProductFee          int64             `json:"product_fee"`            // 商品总价 (单位分)
	SalesMarket         string            `json:"sales_market"`           // 采购市场
	PurchaseOrderLineID int64             `json:"purchase_order_line_id"` // 采购子单id
	CreateTime          int64             `json:"create_time"`            // 创建时间
	Price               string            `json:"price"`                  // 单价
	PurchaseOrderID     int64             `json:"purchase_order_id"`      // 主单id
	SupplierNick        string            `json:"supplier_nick"`          // 供应商昵称
	Currency            string            `json:"currency"`               // 币种
	Properties          *RefundProperties `json:"properties"`             // 商品属性
	ItemImageURL        string            `json:"item_image_url"`         // 商品图片
}

// RefundProperties 商品属性
type RefundProperties struct {
	Options string `json:"options"` // 属性内容
}

// RefundOrderServiceExpr 服务信息 (保留，部分场景返回)
type RefundOrderServiceExpr struct {
	Name  string `json:"name"`  // 名称
	Value string `json:"value"` // 值
}

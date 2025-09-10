package types

// QueryPurchaseOrdersRequest 查询采购订单请求
type QueryPurchaseOrdersRequest struct {
	Status          string  `json:"status,omitempty"`            // 订单状态
	SortType        string  `json:"sort_type,omitempty"`         // ASC/DESC
	PageNo          int     `json:"page_no,omitempty"`           // 页码
	PageSize        int     `json:"page_size,omitempty"`         // 每页数量
	ModifyTimeStart int64   `json:"modify_time_start,omitempty"` // 开始修改时间 (UTC 时间戳)
	ModifyTimeEnd   int64   `json:"modify_time_end,omitempty"`   // 结束修改时间 (UTC 时间戳)
	OuterPurchaseID string  `json:"outer_purchase_id,omitempty"` // ISV 采购单ID
	PurchaseIDS     []int64 `json:"purchase_ids,omitempty"`      // 主单ID列表
}

// QueryPurchaseOrdersResponse 订单明细响应
type QueryPurchaseOrdersResponse struct {
	Success   bool      `json:"success"`    // success flag / 成功标
	ErrorCode string    `json:"error_code"` // error code / 错误码
	ErrorMsg  string    `json:"error_msg"`  // error message / 错误信息
	Data      OrderData `json:"data"`       // 订单明细结果
}

// OrderData 订单数据
type OrderData struct {
	ResultsTotal   int64           `json:"results_total"`   // results total / 总条数
	PageNo         int64           `json:"page_no"`         // page no / 当前页码
	PageSize       int64           `json:"page_size"`       // page size / 每页展示数量
	PurchaseOrders []PurchaseOrder `json:"purchase_orders"` // 采购主单列表明细
}

// PurchaseOrder 采购主单
type PurchaseOrder struct {
	Status                  string             `json:"status"`                          // 订单状态
	OrderSource             string             `json:"order_source"`                    // 订单渠道来源
	Receiver                Receiver           `json:"receiver"`                        // 收件人信息
	PayTime                 string             `json:"pay_time"`                        // 支付成功时间
	PurchaseID              int64              `json:"purchase_id"`                     // 采购单ID
	DistributorNick         string             `json:"distributor_nick"`                // 分销商昵称
	PurchaseCurrency        string             `json:"purchase_currency"`               // 采购金额币种
	OuterPurchaseID         string             `json:"outer_purchase_id"`               // ISV采购单ID
	PurchaseAmount          int64              `json:"purchase_amount"`                 // 采购金额
	NextCloseTimeWithoutPay int64              `json:"next_close_time_without_payment"` // 下一次关单时间
	SubUserID               int64              `json:"sub_user_id"`                     // 子账号id
	OrderRemark             string             `json:"order_remark"`                    // 下单留言 (废弃)
	ProductAmount           int64              `json:"product_amount"`                  // 商品总费用
	SourceMarket            string             `json:"source_market"`                   // 货源市场
	HongbaoDeductionAmount  int64              `json:"hongbao_deduction_amount"`        // 优惠红包金额
	SellerOrderNumber       string             `json:"seller_order_number"`             // ISV单号/平台单号
	ModifyTime              int64              `json:"modify_time"`                     // 采购单修改时间
	PayAmount               string             `json:"pay_amount"`                      // 支付金额
	DomesticPostFee         int64              `json:"domestic_post_fee"`               // 国内运费
	SupplierNick            string             `json:"supplier_nick"`                   // 供应商昵称
	CreatedTime             int64              `json:"created_time"`                    // 采购单创建时间
	PayCurrency             string             `json:"pay_currency"`                    // 实际支付币种
	SubPurchaseOrders       []SubPurchaseOrder `json:"sub_purchase_orders"`             // 子单列表
}

// Receiver 收件人信息
type Receiver struct {
	District string `json:"district"` // 四级地址 （县、区、县级市）
	State    string `json:"state"`    // 二级地址 （省、自治区、直辖市）
	Zip      string `json:"zip"`      // 邮政编码
	Country  string `json:"country"`  // 一级地址 （国家）
	Address  string `json:"address"`  // 详细地址
	City     string `json:"city"`     // 三级地址（市）
}

// SubPurchaseOrder 采购子单
type SubPurchaseOrder struct {
	PayCurrency              string              `json:"pay_currency"`                  // 实际支付币种
	SubPurchaseOrderID       int64               `json:"sub_purchase_order_id"`         // 子单ID
	CloseReason              string              `json:"close_reason"`                  // 关单原因
	SkuID                    int64               `json:"sku_id"`                        // SKU ID
	Status                   string              `json:"status"`                        // 状态
	ErrorCode                string              `json:"error_code"`                    // 订单错误码
	ErrorMessage             string              `json:"error_message"`                 // 订单错误信息
	PurchaseOrderOuterID     string              `json:"purchase_order_outer_id"`       // 淘宝主订单号
	PurchaseOrderLineOuterID string              `json:"purchase_order_line_outer_id"`  // 淘宝子订单号
	LogisticOrders           []LogisticOrder     `json:"logistic_orders"`               // 国内物流信息
	SupplyChainOrders        []SupplyChainOrder  `json:"supply_chain_service_orders"`   // 供应链服务订单
	IntlLogisticOrders       []IntlLogisticOrder `json:"international_logistic_orders"` // 国际物流信息
	CloseTimeWithoutPay      int64               `json:"close_time_without_payment"`    // 子单关单时间
	AlipayPayOrderNo         string              `json:"alipay_pay_order_no"`           // 支付宝流水号
	HongbaoDeductionAmount   string              `json:"hongbao_deduction_amount"`      // 使用红包金额
	DetailErrorCode          string              `json:"detail_error_code"`             // 详细错误码
	ProductAmount            int64               `json:"product_amount"`                // 商品实付价
	OptionFields             OptionFields        `json:"option_fields"`                 // 动态字段
	SkuInfo                  SkuInfo             `json:"sku_info"`                      // SKU信息
	PayTime                  int64               `json:"pay_time"`                      // 支付时间
	OrderRemarkV2            string              `json:"order_remark_v2"`               // 买家备注
	CanConfirmReceipt        bool                `json:"can_confirm_receipt"`           // 是否支持确认收货
	ItemCommissionRate       string              `json:"item_commission_rate"`          // 商品预估佣金率
	RefundID                 int64               `json:"refund_id"`                     // 逆向单Id
	RefundStatus             int64               `json:"refund_status"`                 // 退款状态
	LogisticCompanyName      string              `json:"logistic_company_name"`         // 过期字段: 物流公司
	LogisticNumber           string              `json:"logistic_number"`               // 过期字段: 物流单号
	Quantity                 int64               `json:"quantity"`                      // 商品数量
	RtsTime                  string              `json:"rts_time"`                      // 过期字段: 发货时间
	Title                    string              `json:"title"`                         // 订单标题
	PurchaseCurrency         string              `json:"purchase_currency"`             // 采购币种
	PurchaseAmount           int64               `json:"purchase_amount"`               // 采购价格
	ItemID                   string              `json:"item_id"`                       // 商品ID
	PayAmount                string              `json:"pay_amount"`                    // 实际支付金额
	DomesticPostFee          int64               `json:"domestic_post_fee"`             // 一段物流费用
}

// LogisticOrder 国内物流信息
type LogisticOrder struct {
	LogisticCompanyName string `json:"logistic_company_name"` // 物流公司
	LogisticNumber      string `json:"logistic_number"`       // 物流编号
	RtsTime             int64  `json:"rts_time"`              // 发货时间
}

// IntlLogisticOrder 国际物流信息
type IntlLogisticOrder struct {
	LogisticCompanyName string `json:"logistic_company_name"` // 国际物流公司
	LogisticNumber      string `json:"logistic_number"`       // 国际物流编号
	RtsTime             int64  `json:"rts_time"`              // 国际物流发货时间
}

// SupplyChainOrder 供应链服务订单
type SupplyChainOrder struct {
	ServiceCategory string `json:"service_catrgory"`     // 服务类型
	OrderID         int64  `json:"order_id"`             // 订单ID
	ServiceID       int64  `json:"service_id"`           // 服务ID
	ServiceName     string `json:"service_name"`         // 服务名称
	Status          string `json:"status"`               // 状态
	Price           Price  `json:"price"`                // 价格
	ErrorCode       string `json:"error_code"`           // 错误码
	ErrorMessage    string `json:"error_message"`        // 错误信息
	DetailErrorCode string `json:"detail_error_code"`    // 详细错误码
	DetailErrorMsg  string `json:"detail_error_message"` // 详细错误信息
}

// SkuInfo SKU 信息
type SkuInfo struct {
	SkuPv struct {
		En string `json:"en"` // 英文 SKU 属性 JSON
		Zh string `json:"zh"` // 中文 SKU 属性 JSON
	} `json:"skuPv"`
	SkuImgUrl string `json:"skuImgUrl"` // SKU 图片链接
}

// OptionFields 动态字段
type OptionFields struct {
	SKUImgUrl    string `json:"SKU_IMG_URL"`  // SKU 图片链接
	SupplierNick string `json:"supplierNick"` // 供应商昵称
	En           string `json:"en"`           // 英文属性 JSON
	Zh           string `json:"zh"`           // 中文属性 JSON
	ShopName     string `json:"shopName"`     // 店铺名称
	ShopID       string `json:"shopId"`       // 店铺ID
}

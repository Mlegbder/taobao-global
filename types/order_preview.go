package types

// PurchaseOrderRenderRequest 订单预览请求
type PurchaseOrderRenderRequest struct {
	NeedSupplyChainService bool            `json:"need_supplychain_service"`    // 是否需要供应链服务（跨境物流填 true）
	RenderItemList         []RenderItemReq `json:"-"`                           // 渲染商品列表（JSON字符串）
	WarehouseAddress       *Address        `json:"warehouse_address,omitempty"` // 国内仓库地址（非跨境必填）
	ReceiverAddress        Address         `json:"receiver_address"`            // 收货人地址（跨境必填）
	TaxID                  string          `json:"tax_id,omitempty"`            // 税号（巴西必填）
}

// RenderItemReq 渲染订单商品
type RenderItemReq struct {
	ItemID   string `json:"item_id"`  // 商品ID (mpId)
	SkuID    string `json:"sku_id"`   // SKU ID (mpSkuId)
	Quantity int    `json:"quantity"` // 数量
}

// Address 地址结构体
type Address struct {
	Name        string `json:"name,omitempty"`         // 姓名
	Phone       string `json:"phone,omitempty"`        // 电话
	MobilePhone string `json:"mobile_phone,omitempty"` // 手机
	Country     string `json:"country"`                // 国家（必填）
	State       string `json:"state,omitempty"`        // 省份
	City        string `json:"city,omitempty"`         // 城市
	District    string `json:"district,omitempty"`     // 区/县
	Address     string `json:"address,omitempty"`      // 详细地址
	Zip         string `json:"zip,omitempty"`          // 邮编
}

// PurchaseOrderRenderResponse 订单预览响应
type PurchaseOrderRenderResponse struct {
	Result RenderResult `json:"result"`
}

// RenderResult 结果集
type RenderResult struct {
	InnerErrorMsg string `json:"inner_error_msg"` // 内部错误详细
	Data          Data   `json:"data"`            // 业务数据
	Success       bool   `json:"success"`         // 成功标
	FailItems     string `json:"fail_items"`      // 失败商品
	ErrorCode     string `json:"error_code"`      // 错误码
	ErrorMsg      string `json:"error_msg"`       // 错误原因
}

// Data 业务数据
type Data struct {
	RenderItemList     []RenderItem     `json:"render_item_list"`     // 渲染订单列表
	UnavailableSkuList []UnavailableSku `json:"unavailable_sku_list"` // 不可用商品及原因
	TotalRealPayPrice  Price            `json:"total_real_pay_price"` // 实际支付总价
}

// RenderItem 渲染订单列表项
type RenderItem struct {
	Nick                  string               `json:"nick"`                         // 商家昵称
	ItemPriceInfos        []ItemPriceInfo      `json:"item_price_infos"`             // 商品价格信息列表
	SellerID              int64                `json:"seller_id"`                    // 商家ID
	MainlandShippingFee   Price                `json:"mainland_shipping_fee"`        // 国内物流费用
	OrderFee              Price                `json:"order_fee"`                    // 当前订单费用
	ChooseSupplyChainSvcs []SupplyChainService `json:"choose_supply_chain_services"` // 选择的供应链服务列表
	DispatchPlace         string               `json:"dispatch_place"`               // 发货地
	EstimatedDeliveryTime string               `json:"estimated_delivery_time"`      // 物流时效
	CurrencyCode          string               `json:"currency_code"`                // 币种
	DeliveryFee           int64                `json:"delivery_fee"`                 // 物流费用
}

// ItemPriceInfo 商品价格信息
type ItemPriceInfo struct {
	ItemID        int64 `json:"item_id"`        // 分销商品id
	OriginPrice   Price `json:"origin_price"`   // 商品单价
	SkuID         int64 `json:"sku_id"`         // 分销商品skuid
	Quantity      int64 `json:"quantity"`       // 商品数量
	DiscountPrice Price `json:"discount_price"` // 折扣后金额小记
}

// Price 金额结构
type Price struct {
	Amount   int64  `json:"amount"`   // 金额（分）
	Currency string `json:"currency"` // 币种
}

// SupplyChainService 供应链服务
type SupplyChainService struct {
	ShippingFee     Price        `json:"shipping_fee"`     // 运费/服务费
	Name            string       `json:"name"`             // 方案名称
	Description     string       `json:"description"`      // 描述
	OptionID        int64        `json:"option_id"`        // 服务方案id
	IsMustSelect    bool         `json:"is_must_select"`   // 是否必选
	ShippingTime    ShippingTime `json:"shipping_time"`    // 物流方案时间
	ServiceCategory string       `json:"service_category"` // 服务类目
}

// ShippingTime 物流时间
type ShippingTime struct {
	Mode int64 `json:"mode"` // 1 工作日 2 自然天
	Min  int64 `json:"min"`  // 最短
	Max  int64 `json:"max"`  // 最长
}

// UnavailableSku 不可用商品信息
type UnavailableSku struct {
	ItemID    int64  `json:"item_id"`    // 分销商品id
	SkuID     int64  `json:"sku_id"`     // 分销商品skuid
	Reason    string `json:"reason"`     // 失败原因
	ErrorCode string `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误原因
}

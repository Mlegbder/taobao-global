package types

// PurchaseOrderRenderRequest 订单预览请求
type PurchaseOrderRenderRequest struct {
	NeedSupplyChainService bool     `json:"need_supplychain_service"`    // 是否需要供应链服务（跨境物流填 true）
	RenderItemList         string   `json:"render_item_list"`            // 渲染商品列表（JSON字符串）
	WarehouseAddress       *Address `json:"warehouse_address,omitempty"` // 国内仓库地址（非跨境必填）
	ReceiverAddress        Address  `json:"receiver_address"`            // 收货人地址（跨境必填）
	TaxID                  string   `json:"tax_id,omitempty"`            // 税号（巴西必填）
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
	Result        RenderResult `json:"result"`
	InnerErrorMsg string       `json:"inner_error_msg"`
	Data          RenderData   `json:"data"`
}

// RenderResult 渲染结果
type RenderResult struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	FailItems string `json:"fail_items"`
}

// RenderData 订单预览数据
type RenderData struct {
	RenderItemList       []RenderItem         `json:"render_item_list"`
	CurrencyCode         string               `json:"currency_code"`
	DeliveryFee          int64                `json:"delivery_fee"`
	SellerID             int64                `json:"seller_id"`
	MainlandShippingFee  *Money               `json:"mainland_shipping_fee"`
	OrderFee             *Money               `json:"order_fee"`
	ChooseSupplyServices []SupplyChainService `json:"choose_supply_chain_services"`
	UnavailableSkuList   []UnavailableSku     `json:"unavailable_sku_list"`
	TotalRealPayPrice    *Money               `json:"total_real_pay_price"`
}

// RenderItem 渲染后的商品条目
type RenderItem struct {
	Nick                  string      `json:"nick"`
	ItemPriceInfos        []ItemPrice `json:"item_price_infos"`
	DispatchPlace         string      `json:"dispatch_place"`
	EstimatedDeliveryTime string      `json:"estimated_delivery_time"`
}

// ItemPrice 商品价格信息
type ItemPrice struct {
	ItemID        int64  `json:"item_id"`
	OriginPrice   *Money `json:"origin_price"`
	Amount        *Money `json:"amount"`
	Currency      string `json:"currency"`
	SkuID         int64  `json:"sku_id"`
	Quantity      int    `json:"quantity"`
	DiscountPrice *Money `json:"discount_price"`
}

// Money 通用金额结构
type Money struct {
	Amount   int64  `json:"amount"` // 单位分
	Currency string `json:"currency"`
}

// SupplyChainService 供应链服务
type SupplyChainService struct {
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	OptionID        int64         `json:"option_id"`
	IsMustSelect    bool          `json:"is_must_select"`
	ServiceCategory string        `json:"service_category"`
	ShippingFee     *Money        `json:"shipping_fee"`
	ShippingTime    *ShippingTime `json:"shipping_time"`
}

// ShippingTime 物流时效
type ShippingTime struct {
	Mode int `json:"mode"`
	Min  int `json:"min"`
	Max  int `json:"max"`
}

// UnavailableSku 不可用SKU信息
type UnavailableSku struct {
	ItemID int64  `json:"item_id"`
	SkuID  int64  `json:"sku_id"`
	Reason string `json:"reason"`
}

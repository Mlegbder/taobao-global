package types

// ItemDetailRequest 获取商品详情请求
type ItemDetailRequest struct {
	ItemResource string   `json:"item_resource"`          // 必须：商品来源 taobao
	ItemID       string   `json:"item_id"`                // 必须：淘宝/天猫商品ID
	IncludeTags  []string `json:"include_tags,omitempty"` // 可选：待匹配标签
	Language     string   `json:"language,omitempty"`     // 可选：翻译语种 en|vi|ru|ko|ja
}

// ItemDetailResponse 获取商品详情响应
type ItemDetailResponse struct {
	BizErrorCode string     `json:"biz_error_code"` // 错误码
	BizErrorMsg  string     `json:"biz_error_msg"`  // 错误信息
	Data         ItemDetail `json:"data"`           // 商品详情
}

// ItemDetail 商品详情
type ItemDetail struct {
	ItemResource      string              `json:"item_resource"`       // 来源 taobao
	ItemID            int64               `json:"item_id"`             // 商品ID
	Title             string              `json:"title"`               // 标题
	PicUrls           []string            `json:"pic_urls"`            // 主图
	ShopName          string              `json:"shop_name"`           // 店铺名称
	Price             int64               `json:"price"`               // 原价 (单位: 分)
	PromotionPrice    int64               `json:"promotion_price"`     // 优惠价 (单位: 分)
	Description       string              `json:"description"`         // 商品详情
	Properties        []ItemProperty      `json:"properties"`          // 商品参数信息
	SkuList           []SkuItem           `json:"sku_list"`            // SKU 列表
	CouponPrice       int64               `json:"coupon_price"`        // 券后价 (单位: 分)
	ShopID            int64               `json:"shop_id"`             // 店铺ID
	PropertyImages    []PropertyImage     `json:"property_image_list"` // 属性图片
	Tags              []string            `json:"tags"`                // 确认标签
	MultiLangInfo     MultiLangInfoDetail `json:"multi_language_info"` // 多语言信息
	PromotionDisplays []PromotionDisplay  `json:"promotion_displays"`  // 优惠列表
}

// ItemProperty 商品属性
type ItemProperty struct {
	PropID    int64  `json:"prop_id"`    // 属性ID
	PropName  string `json:"prop_name"`  // 属性名称
	ValueID   int64  `json:"value_id"`   // 属性值ID
	ValueName string `json:"value_name"` // 属性值
	// ValueDesc 在多语言场景下会出现 (属性值+备注)，这里省略了，可在 MultiLangInfo 里定义
}

// SkuItem 商品SKU
type SkuItem struct {
	SkuID          int64          `json:"sku_id"`          // SKU ID
	Quantity       int            `json:"quantity"`        // 库存数量
	Price          int64          `json:"price"`           // 原价 (单位: 分)
	PromotionPrice int64          `json:"promotion_price"` // 优惠价 (单位: 分)
	PicURL         string         `json:"pic_url"`         // 图片链接
	Properties     []ItemProperty `json:"properties"`      // SKU属性列表
	CouponPrice    int64          `json:"coupon_price"`    // 券后价 (单位: 分)
}

// PropertyImage 属性图片
type PropertyImage struct {
	Properties string `json:"properties"` // 属性 (如颜色:红色)
	ImageURL   string `json:"image_url"`  // 图片链接
}

// MultiLangInfoDetail 多语言翻译信息
type MultiLangInfoDetail struct {
	Title        string            `json:"title"`          // 多语言标题
	Properties   []ItemProperty    `json:"properties"`     // 多语言属性
	SkuProps     []SkuLangProperty `json:"sku_properties"` // 多语言SKU属性
	Language     string            `json:"language"`       // 翻译语种
	MainImageURL string            `json:"main_image_url"` // 多语言主图
	CouponPrice  int64             `json:"coupon_price"`   // 多语言券后价 (单位: 分)
}

// SkuLangProperty 多语言SKU属性
// SkuLangProperty 多语言SKU属性
type SkuLangProperty struct {
	SkuID      int64          `json:"sku_id"`     // SKU ID
	Properties []ItemProperty `json:"properties"` // 多语言SKU属性列表
}

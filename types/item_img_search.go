package types

// ImgSearchRequest 图片搜索请求
type ImgSearchRequest struct {
	PicURL      string   `json:"pic_url,omitempty"`      // 可选: 图片链接
	IncludeTags []string `json:"include_tags,omitempty"` // 可选: 活动标签
	Language    string   `json:"language,omitempty"`     // 可选: 翻译语种 en|vi|ru|ko|ja
	ImageID     string   `json:"image_id,omitempty"`     // 可选: 图片ID (通过上传接口获得)
}

// ImgSearchResponse 图片搜索响应
type ImgSearchResponse struct {
	BizErrorCode string          `json:"biz_error_code"` // 错误码
	BizErrorMsg  string          `json:"biz_error_msg"`  // 错误原因
	Data         []ImgSearchItem `json:"data"`           // 搜索结果 (最多100条)
}

// ImgSearchItem 图搜结果条目
type ImgSearchItem struct {
	ItemID       int64              `json:"item_id"`
	Title        string             `json:"title"`
	MainImageURL string             `json:"main_image_url"`
	Price        string             `json:"price"` // 单位: 元 (字符串类型)
	Inventory    int64              `json:"inventory"`
	ShopName     string             `json:"shop_name"`
	Tags         []string           `json:"tags"`
	MultiLang    *MultiLangInfo     `json:"multi_language_info"`
	CouponPrice  string             `json:"coupon_price"`
	Promotions   []PromotionDisplay `json:"promotion_displays"`
}

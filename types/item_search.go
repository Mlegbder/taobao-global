package types

// ItemSearchRequest 商品搜索请求参数
type ItemSearchRequest struct {
	Keyword     string   `json:"keyword,omitempty"`      // 搜索关键字
	IncludeTags []string `json:"include_tags,omitempty"` // 活动标签，例如 activity_202311_1_tb_manjian
	Sort        string   `json:"sort,omitempty"`         // 排序方式: PRICE_ASC, PRICE_DESC, SALE_QTY_ASC, SALE_QTY_DESC
	PageNo      int      `json:"page_no,omitempty"`      // 页码（默认 1）
	PageSize    int      `json:"page_size,omitempty"`    // 页大小（默认 10，上限 20）
	Filters     []string `json:"filters,omitempty"`      // 筛选项，例如 min_price:2000, max_price:10000 (单位: 分)
	Language    string   `json:"language,omitempty"`     // 翻译语种 en|vi|ru|ko|ja
	ShopID      int      `json:"shop_id,omitempty"`      // 店铺ID
}

// ItemSearchResponse 商品搜索响应
type ItemSearchResponse struct {
	BizErrorCode string   `json:"biz_error_code"` // 错误码
	BizErrorMsg  string   `json:"biz_error_msg"`  // 错误原因
	Data         ItemPage `json:"data"`           // 搜索结果
}

// ItemPage 分页信息
type ItemPage struct {
	PageNo   int         `json:"page_no"`   // 当前页码
	PageSize int         `json:"page_size"` // 每页大小
	Items    []ItemEntry `json:"data"`      // 商品列表
}

// ItemEntry 单个商品信息
type ItemEntry struct {
	ItemID           int64              `json:"item_id"`             // 商品ID
	MainImageURL     string             `json:"main_image_url"`      // 商品主图
	Price            string             `json:"price"`               // 优惠价（单位: 元）
	Inventory        int                `json:"inventory"`           // 库存
	Tags             []string           `json:"tags"`                // 活动标签
	ShopName         string             `json:"shop_name"`           // 店铺名称
	Title            string             `json:"title"`               // 商品标题
	MultiLangInfo    MultiLangInfo      `json:"multi_language_info"` // 多语言信息
	CouponPrice      string             `json:"coupon_price"`        // 券后价（仅供参考）
	PromotionDisplay []PromotionDisplay `json:"promotion_displays"`  // 营销优惠
}

// MultiLangInfo 商品翻译信息
type MultiLangInfo struct {
	Language     string `json:"language"`       // 翻译语种
	Title        string `json:"title"`          // 翻译标题
	MainImageURL string `json:"main_image_url"` // 翻译主图
}

// PromotionDisplay 商品优惠展示
type PromotionDisplay struct {
	TypeName          string          `json:"type_name"`           // 优惠类型名称
	PromotionInfoList []PromotionInfo `json:"promotion_info_list"` // 优惠明细
}

// PromotionInfo 优惠详情
type PromotionInfo struct {
	ActivityCode  string `json:"activity_code"`  // 活动编码
	PromotionName string `json:"promotion_name"` // 优惠名称
}

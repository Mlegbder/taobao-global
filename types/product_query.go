package types

// QueryAllProductRequest 获取商品请求
type QueryAllProductRequest struct {
	ItemID           string `json:"item_id"`                      // 必填: 淘宝商品id
	ItemSourceMarket string `json:"item_source_market,omitempty"` // 可选: CBU_MARKET (1688)，默认淘宝天猫
}

// QueryAllProductResponse 获取商品响应
type QueryAllProductResponse struct {
	Success   bool             `json:"success"`
	FailItems string           `json:"fail_items"`
	ErrorCode string           `json:"error_code"`
	ErrorMsg  string           `json:"error_msg"`
	Data      *ProductInfoData `json:"data"`
}

// ProductInfoData 商品信息
type ProductInfoData struct {
	Quantity       int64              `json:"quantity"`
	CategoryPath   string             `json:"category_path"`
	ShopName       string             `json:"shop_name"`
	Description    string             `json:"description"`
	ProductUnit    string             `json:"product_unit"`
	PicURLs        []string           `json:"pic_urls"`
	Title          string             `json:"title"`
	MpID           string             `json:"mp_id"`
	CategoryName   string             `json:"category_name"`
	ItemID         int64              `json:"item_id"`
	UserNick       string             `json:"user_nick"`
	Price          int64              `json:"price"`
	BeginAmount    int64              `json:"begin_amount"`
	Status         string             `json:"status"`
	SKUList        []ProductSKU       `json:"sku_list"`
	PromotionPrice int64              `json:"promotion_price"`
	CouponPrice    int64              `json:"coupon_price"`
	PostFee        int64              `json:"postFee"`
	ShopID         int64              `json:"shop_id"`
	CategoryID     string             `json:"category_id"`
	ItemType       string             `json:"item_type"`
	Promotions     []PromotionDisplay `json:"promotion_displays"`
}

// ProductSKU SKU 信息
type ProductSKU struct {
	PicURL         string            `json:"pic_url"`
	Quantity       int64             `json:"quantity"`
	Price          int64             `json:"price"`
	SkuID          int64             `json:"sku_id"`
	Properties     []ProductProperty `json:"properties"`
	Status         string            `json:"status"`
	PromotionPrice int64             `json:"promotion_price"`
	CouponPrice    int64             `json:"coupon_price"`
	MpSkuID        int64             `json:"mp_skuId"`
	PostFee        int64             `json:"postFee"`
}

type ProductProperty struct {
	ValueID   int64  `json:"value_id"`   // 属性值id
	ValueName string `json:"value_name"` // 属性值
	ValueDesc string `json:"value_desc"` // 属性值+备注
	PropID    int64  `json:"prop_id"`    // 属性id
	PropName  string `json:"prop_name"`  // 属性名称
}

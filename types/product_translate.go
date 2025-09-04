package types

// ProductTranslateRequest 商品翻译请求
type ProductTranslateRequest struct {
	ItemID   string `json:"item_id"`            // 必填: 铺货商品id (mp_id)
	Language string `json:"language,omitempty"` // 可选: en/ru/ko/ja/vi (默认中文)
}

// ProductTranslateResponse 商品翻译响应
type ProductTranslateResponse struct {
	Result        *TranslateResult      `json:"result"`
	InnerErrorMsg string                `json:"inner_error_msg"`
	Data          *ProductTranslateData `json:"data"`
	Success       bool                  `json:"success"`
	FailItems     string                `json:"fail_items"`
	ErrorCode     string                `json:"error_code"`
	ErrorMsg      string                `json:"error_msg"`
}

// TranslateResult 通用结果
type TranslateResult struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// ProductTranslateData 商品翻译数据
type ProductTranslateData struct {
	ItemID        string                 `json:"item_id"`
	MainImage     string                 `json:"main_image"`
	Language      string                 `json:"language"`
	Title         string                 `json:"title"`
	Properties    []ProductProperty      `json:"properties"`
	SkuProperties []SkuTranslateProperty `json:"sku_properties"`
}

// ProductProperty 普通属性
type ProductProperty struct {
	PropID    int64  `json:"prop_id"`
	PropName  string `json:"prop_name"`
	ValueID   int64  `json:"value_id"`
	ValueName string `json:"value_name"`
	ValueDesc string `json:"value_desc"`
}

// SkuTranslateProperty SKU 属性
type SkuTranslateProperty struct {
	SkuID      int64             `json:"sku_id"`
	Properties []ProductProperty `json:"properties"`
}

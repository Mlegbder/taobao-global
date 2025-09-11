package types

// ProductTranslateRequest 商品翻译请求
type ProductTranslateRequest struct {
	ItemID   string `json:"item_id"`            // 必填: 铺货商品id (mp_id)
	Language string `json:"language,omitempty"` // 可选: en/ru/ko/ja/vi (默认中文)
}

// ProductInfoTranResponse API: 商品信息翻译
type ProductInfoTranResponse struct {
	Result ProductInfoTranResult `json:"result"` // 结果集
}

// ProductInfoTranResult 结果集
type ProductInfoTranResult struct {
	InnerErrorMsg string              `json:"inner_error_msg"` // 内部错误详情
	Data          ProductInfoTranData `json:"data"`            // 业务结果
	Success       bool                `json:"success"`         // 成功标
	FailItems     string              `json:"fail_items"`      // 失败条目
	ErrorCode     string              `json:"error_code"`      // 错误码
	ErrorMsg      string              `json:"error_msg"`       // 错误信息
}

// ProductInfoTranData 业务结果
type ProductInfoTranData struct {
	SkuProperties []TranSkuProperty `json:"sku_properties"` // sku属性对象
	ItemID        string            `json:"item_id"`        // 铺货商品id
	MainImage     string            `json:"main_image"`     // 商品主图
	Language      string            `json:"language"`       // 返回的语种信息
	Title         string            `json:"title"`          // 商品标题
	Properties    []TranProperty    `json:"properties"`     // 属性对象
}

// TranSkuProperty sku属性对象
type TranSkuProperty struct {
	SkuID      int64          `json:"sku_id"`     // skuId
	Properties []TranProperty `json:"properties"` // 属性对象
}

// TranProperty 属性对象
type TranProperty struct {
	ValueID   int64  `json:"value_id"`   // 属性值id
	ValueName string `json:"value_name"` // 属性值
	ValueDesc string `json:"value_desc"` // 属性值+备注
	PropID    int64  `json:"prop_id"`    // 属性id
	PropName  string `json:"prop_name"`  // 属性名称
}

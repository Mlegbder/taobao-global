package types

// ImageUploadRequest 图片上传请求
type ImageUploadRequest struct {
	ImageBase64 string `json:"image_base64"` // 必填: Base64 编码的 JPG/PNG/WEBP 图片，最大 3MB
}

// ImageUploadResponse 图片上传响应
type ImageUploadResponse struct {
	BizErrorCode string           `json:"biz_error_code"` // 错误码
	BizErrorMsg  string           `json:"biz_error_msg"`  // 错误原因
	Data         *ImageUploadData `json:"data"`
}

// ImageUploadData 图片存储信息
type ImageUploadData struct {
	ImageID string `json:"image_id"` // 图片ID (主要用于图搜)
}

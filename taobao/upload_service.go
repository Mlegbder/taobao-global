package taobao

import (
	"encoding/json"
	"taobao-global/consts"
	"taobao-global/types"
	"taobao-global/utils"
)

// UploadService 提供上传相关API
type UploadService struct {
	client *Client
}

// Image 图片上传
func (s *UploadService) Image(req types.ImageUploadRequest, accessToken string) (*types.ImageUploadResponse, error) {
	params := map[string]string{
		"access_token": accessToken,
		"image_base64": req.ImageBase64,
	}

	baseConf := s.client.Base
	baseConf.ApiEndpoint = consts.TaoBaoApiImageUpload

	respBytes, err := utils.Execute(params, baseConf)
	if err != nil {
		return nil, err
	}

	var resp types.ImageUploadResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

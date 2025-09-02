package types

// TokenRequest 请求参数
type TokenRequest struct {
	Code string `json:"code"` // 必须: oauth code, 从回调URL获取
}

// TokenResponse 响应参数
type TokenResponse struct {
	ExpiresIn        int64  `json:"expires_in"`         // Access Token 过期时间（秒）
	AccountID        string `json:"account_id"`         // Account ID (可能为 null)
	SellerID         string `json:"seller_id"`          // 卖家 ID
	UserID           string `json:"user_id"`            // 用户 ID
	ShortCode        string `json:"short_code"`         // 卖家 short code
	AccountPlatform  string `json:"account_platform"`   // 账户平台
	AccessToken      string `json:"access_token"`       // 访问 token
	Account          string `json:"account"`            // 登录用户账户
	RefreshToken     string `json:"refresh_token"`      // 刷新 token
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // 刷新 token 过期时间（秒）
}

// RefreshTokenRequest 请求参数
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"` // 必须: refresh_token
}

// RefreshTokenResponse 响应参数
type RefreshTokenResponse struct {
	ExpiresIn        int64   `json:"expires_in"`         // Access Token 过期时间（秒）
	AccountID        *string `json:"account_id"`         // Account ID（可能为 null）
	UserID           string  `json:"user_Id"`            // 用户 ID
	SellerID         string  `json:"seller_Id"`          // 卖家 ID
	ShortCode        string  `json:"short_code"`         // 卖家短码
	AccountPlatform  string  `json:"account_platform"`   // 账户平台
	AccessToken      string  `json:"access_token"`       // 新的访问 token
	Account          string  `json:"account"`            // 登录用户账户
	RefreshToken     string  `json:"refresh_token"`      // 刷新 token
	RefreshExpiresIn int64   `json:"refresh_expires_in"` // 刷新 token 过期时间（秒）
	ErrorCode        string  `json:"error_code"`         // 错误码（失败时返回）
}

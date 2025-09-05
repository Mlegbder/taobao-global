package types

type TaobaoBase struct {
	ApiEndpoint  string
	AppKey       string
	AppSecret    string
	Api          string
	AccessToken  string
	RefreshToken string
}

type BaseResponse struct {
	Data      interface{} `json:"data"`       // 业务数据，不固定，具体接口定义
	Code      string      `json:"code"`       // 响应码 (通常 "0" 表示成功)
	RequestID string      `json:"request_id"` // 请求ID
	TraceID   string      `json:"_trace_id_"` // Trace链路ID
}

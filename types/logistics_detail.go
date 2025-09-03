package types

// GetLogisticsDetailRequest 获取子单物流信息请求
type GetLogisticsDetailRequest struct {
	PurchaseOrderLineID int64 `json:"purchase_order_line_id"` // 必填: 采购子单id
}

// GetLogisticsDetailResponse 获取子单物流信息响应
type GetLogisticsDetailResponse struct {
	Result        *LogisticsResult `json:"result"`
	InnerErrorMsg string           `json:"inner_error_msg"`
	Data          *LogisticsData   `json:"data"`
	Success       bool             `json:"success"`
	FailItems     string           `json:"fail_items"`
	ErrorCode     string           `json:"error_code"`
	ErrorMsg      string           `json:"error_msg"`
}

// LogisticsResult 通用结果
type LogisticsResult struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// LogisticsData 物流数据
type LogisticsData struct {
	PnmLogisticsDetails  []LogisticsDetail `json:"pnm_logistics_details"`
	PurchaseOrderOuterID int64             `json:"purchase_order_outer_id"`
	LogisticsStatus      string            `json:"logistics_status"`
	LogisticsDesc        string            `json:"logistics_desc"`
}

// LogisticsDetail 包裹物流详情
type LogisticsDetail struct {
	MailNo           string             `json:"mail_no"`
	LogisticsTraces  []LogisticsTrace   `json:"logistics_traces"`
	LogisticsPartner []LogisticsPartner `json:"logistics_partner"`
	Receiver         *LogisticsReceiver `json:"receiver"`
	LogisticsGoods   []LogisticsGoods   `json:"logistics_goods"`
}

// LogisticsTrace 物流轨迹
type LogisticsTrace struct {
	NextCity     string `json:"next_city"`
	StatusDesc   string `json:"status_desc"`
	Province     string `json:"province"`
	City         string `json:"city"`
	NextNodeName string `json:"next_node_name"`
	District     string `json:"district"`
	Action       string `json:"action"`
	FacilityName string `json:"facility_name"`
	Time         string `json:"time"`
	StanderdDesc string `json:"standerd_desc"`
	Status       string `json:"status"`
}

// LogisticsPartner 物流公司
type LogisticsPartner struct {
	CompanyImageURL string `json:"company_image_url"`
	Company         string `json:"company"`
	CompanyContact  string `json:"company_contact"`
}

// LogisticsReceiver 收件人信息
type LogisticsReceiver struct {
	ZipCode  string `json:"zip_code"`
	Address  string `json:"address"`
	Province string `json:"province"`
	City     string `json:"city"`
	Phone    string `json:"phone"`
	District string `json:"district"`
	Name     string `json:"name"`
}

// LogisticsGoods 包裹商品
type LogisticsGoods struct {
	Quantity    int64  `json:"quantity"`
	ImageURL    string `json:"image_url"`
	Property    string `json:"property"`
	Title       string `json:"title"`
	GoodsNumber int64  `json:"goods_number"`
}

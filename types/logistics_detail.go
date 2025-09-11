package types

// GetLogisticsDetailRequest 获取子单物流信息请求
type GetLogisticsDetailRequest struct {
	PurchaseOrderLineID int64 `json:"purchase_order_line_id"` // 必填: 采购子单id
}

// LogisticsDetailResponse API: 物流详情
type LogisticsDetailResponse struct {
	Result LogisticsDetailResult `json:"result"` // 结果集
}

// LogisticsDetailResult 结果集
type LogisticsDetailResult struct {
	InnerErrorMsg string              `json:"inner_error_msg"` // 内部错误详情
	Data          LogisticsDetailData `json:"data"`            // 业务结果
	Success       bool                `json:"success"`         // 成功标
	FailItems     string              `json:"fail_items"`      // 失败条目
	ErrorCode     string              `json:"error_code"`      // 错误码
	ErrorMsg      string              `json:"error_msg"`       // 错误信息
}

// LogisticsDetailData 业务结果
type LogisticsDetailData struct {
	PnmLogisticsDetails  []PnmLogisticsDetail `json:"pnm_logistics_details"`   // 物流包裹清单
	PurchaseOrderOuterID int64                `json:"purchase_order_outer_id"` // 外部采购单ID
}

// PnmLogisticsDetail 物流包裹清单
type PnmLogisticsDetail struct {
	MailNo           string             `json:"mail_no"`           // 运单号
	LogisticsTraces  []LogisticsTrace   `json:"logistics_traces"`  // 物流轨迹信息
	LogisticsPartner []LogisticsPartner `json:"logistics_partner"` // 包裹运输公司
	Receiver         LogisticsReceiver  `json:"receiver"`          // 收件人信息
	LogisticsGoods   []LogisticsGoods   `json:"logistics_goods"`   // 包裹商品信息
	GoodsNumber      int                `json:"goods_number"`      // 包裹中商品数量
	LogisticsStatus  string             `json:"logistics_status"`  // 当前物流状态code 示例 DELIVERING
	LogisticsDesc    string             `json:"logistics_desc"`    // 当前物流状态描述 示例 派送中
}

// LogisticsTrace 物流轨迹信息
type LogisticsTrace struct {
	NextCity     string `json:"next_city"`      // 下一站城市
	StatusDesc   string `json:"status_desc"`    // 节点状态描述 示例 派送中
	Province     string `json:"province"`       // 事件发生省份
	City         string `json:"city"`           // 事件发生城市
	NextNodeName string `json:"next_node_name"` // 下一个操作节点名称
	District     string `json:"district"`       // 事件发生区域
	Action       string `json:"action"`         // 动作类型
	FacilityName string `json:"facility_name"`  // 网点/中转站/仓库
	Time         string `json:"time"`           // 节点发生时间
	StanderdDesc string `json:"standerd_desc"`  // 标准化文案
	Status       string `json:"status"`         // 节点状态code 示例 DELIVERING
}

// LogisticsPartner 包裹运输公司
type LogisticsPartner struct {
	CompanyImageURL string `json:"company_image_url"` // 物流公司logo 图片
	Company         string `json:"company"`           // 物流公司名字
	CompanyContact  string `json:"company_contact"`   // 物流公司联系方式
}

// LogisticsReceiver 收件人信息
type LogisticsReceiver struct {
	ZipCode  string `json:"zip_code"` // 邮编
	Address  string `json:"address"`  // 详细地址
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	Phone    string `json:"phone"`    // 联系方式
	District string `json:"district"` // 区
	Name     string `json:"name"`     // 名字
}

// LogisticsGoods 包裹商品信息
type LogisticsGoods struct {
	Quantity int    `json:"quantity"`  // 数量
	ImageURL string `json:"image_url"` // 商品图片链接
	Property string `json:"property"`  // 商品销售属性
	Title    string `json:"title"`     // 商品名称
}

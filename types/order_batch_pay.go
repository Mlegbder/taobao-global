package types

// BatchPayPurchaseOrderRequest 批量支付采购单请求
type BatchPayPurchaseOrderRequest struct {
	PurchaseOrderIDList []int64 `json:"purchaseOrderIdList"` // 必填: 采购主单ID列表（最多10个）
}

// BatchPayPurchaseOrderResponse 批量支付采购单响应
type BatchPayPurchaseOrderResponse struct {
	Code      string          `json:"code"`       // 响应状态码
	Data      *BatchPayResult `json:"data"`       // 结果数据
	Success   bool            `json:"success"`    // 是否成功（注意返回的是字符串 "true"/"false"）
	ErrorCode string          `json:"error_code"` // 错误码
	ErrorMsg  string          `json:"error_msg"`  // 错误信息
	RequestID string          `json:"request_id"` // 请求 ID
}

// BatchPayResult 批量支付结果
type BatchPayResult struct {
	PayFailurePurchaseOrderIDs []string        `json:"pay_failure_purchase_order_ids"` // 支付失败的主单ID
	WillPayPurchaseOrderIDs    []string        `json:"will_pay_purchase_order_ids"`    // 待支付的主单ID
	PayFailedResults           []PayFailedInfo `json:"pay_failed_results"`             // 支付失败明细
}

// PayFailedInfo 支付失败明细
type PayFailedInfo struct {
	GspOrderID   int64  `json:"gspOrderId"`   // GSP订单ID
	ErrorMessage string `json:"errorMessage"` // 错误信息
	ErrorCode    string `json:"errorCode"`    // 错误码
	Class        string `json:"class"`        // Java类路径
}

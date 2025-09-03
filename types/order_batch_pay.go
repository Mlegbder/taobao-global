package types

// BatchPayPurchaseOrderRequest 批量支付采购单请求
type BatchPayPurchaseOrderRequest struct {
	PurchaseOrderIDList []int64 `json:"purchaseOrderIdList"` // 必填: 采购主单ID列表（最多10个）
}

// BatchPayPurchaseOrderResponse 批量支付采购单响应
type BatchPayPurchaseOrderResponse struct {
	Data      *BatchPayResult `json:"data"`
	Success   bool            `json:"success"`    // 是否成功
	ErrorCode string          `json:"error_code"` // 错误码
	ErrorMsg  string          `json:"error_msg"`  // 错误信息
}

// BatchPayResult 批量支付结果
type BatchPayResult struct {
	PayFailurePurchaseOrderIDs []int64         `json:"pay_failure_purchase_order_ids"` // 支付失败的主单ID
	WillPayPurchaseOrderIDs    []int64         `json:"will_pay_purchase_order_ids"`    // 待支付的主单ID
	PayFailedResults           []PayFailedInfo `json:"pay_failed_results"`             // 支付失败明细
}

// PayFailedInfo 支付失败明细
type PayFailedInfo struct {
	PurchaseOrderID int64  `json:"purchase_order_id"`
	Reason          string `json:"reason"`
}

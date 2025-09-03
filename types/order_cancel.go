package types

// AsynCancelPurchaseOrderRequest 异步取消采购单请求
type AsynCancelPurchaseOrderRequest struct {
	PurchaseID          string   `json:"purchase_id"`                         // 必填: 主采购单号
	SubPurchaseOrderIDs []string `json:"sub_purchase_orderId_list,omitempty"` // 可选: 子单号列表
	CancelReason        string   `json:"cancel_reason"`                       // 必填: 取消原因
	CancelRemark        string   `json:"cancel_remark,omitempty"`             // 可选: 备注
}

// AsynCancelPurchaseOrderResponse 异步取消采购单响应
type AsynCancelPurchaseOrderResponse struct {
	Success   bool   `json:"success"`    // 是否成功 (发起任务成功)
	ErrorCode string `json:"error_code"` // 错误码
	ErrorMsg  string `json:"error_msg"`  // 错误原因
}

package consts

// PurchaseOrderStatus 采购单状态常量
const (
	// BULIDING 采购单创建中
	PurchaseOrderStatusBuilding = "BULIDING"

	// WAIT_BUYER_P 等待付款
	PurchaseOrderStatusWaitBuyerPay = "WAIT_BUYER_P"

	// WAIT_SELLER_SEND_GOODS 已付款，待发货
	PurchaseOrderStatusWaitSellerSend = "WAIT_SELLER_SEND_GOODS"

	// WAIT_BUYER_CONFIRM_GOODS 已付款，已发货
	PurchaseOrderStatusWaitBuyerConfirm = "WAIT_BUYER_CONFIRM_GOODS"

	// TRADE_CLOSED 交易关闭
	PurchaseOrderStatusClosed = "TRADE_CLOSED"
)

// PurchaseOrderStatusDesc 状态对应的中文描述
var PurchaseOrderStatusDesc = map[string]string{
	PurchaseOrderStatusBuilding:         "采购单创建中",
	PurchaseOrderStatusWaitBuyerPay:     "等待付款",
	PurchaseOrderStatusWaitSellerSend:   "已付款，待发货",
	PurchaseOrderStatusWaitBuyerConfirm: "已付款，已发货",
	PurchaseOrderStatusClosed:           "交易关闭",
}

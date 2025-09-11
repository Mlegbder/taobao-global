package consts

// GoodsStatus represents goods delivery states
const (
	GoodsStatusNotShipped      = 1 // 未发货 / Not shipped
	GoodsStatusShipped         = 2 // 已发货 / Shipped
	GoodsStatusNoGoodsReceived = 3 // 未收到货 / No goods received
	GoodsStatusReceived        = 4 // 已收到货 / Received the goods
	GoodsStatusSentBack        = 5 // 已寄回 / Sent back
	GoodsStatusSellerConfirmed = 6 // 卖家确认收货 / Seller's confirmation of receipt
)

// GoodsStatusText maps goods status codes to their descriptions (Chinese + English)
var GoodsStatusText = map[int]string{
	GoodsStatusNotShipped:      "未发货 / Not shipped",
	GoodsStatusShipped:         "已发货 / Shipped",
	GoodsStatusNoGoodsReceived: "未收到货 / No goods received",
	GoodsStatusReceived:        "已收到货 / Received the goods",
	GoodsStatusSentBack:        "已寄回 / Sent back",
	GoodsStatusSellerConfirmed: "卖家确认收货 / Seller's confirmation of receipt",
}

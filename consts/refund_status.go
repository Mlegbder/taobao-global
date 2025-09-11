package consts

// RefundStatus represents refund process states
const (
	RefundStatusNoApply       = 0   // 没有申请退款 / No apply refund
	RefundStatusBuyerApplied  = 10  // 买家已经申请退款，等待卖家同意 / Buyer has applied for a refund, waiting for the seller to agree
	RefundStatusSellerAgreed  = 20  // 卖家已经同意退款，等待买家退货 / The seller has agreed to refund, waiting for the buyer to return the goods
	RefundStatusBuyerReturned = 30  // 买家已经退货，等待卖家确认收货 / Buyer has returned the goods, waiting for the seller to confirm receipt
	RefundStatusSuccess       = 100 // 退款成功 / Successful refund
	RefundStatusSellerRefused = -10 // 卖家拒绝退款 / The seller refused to refund
	RefundStatusClosed        = -20 // 退款关闭 / Refund Close
)

var RefundStatusMap = map[int]string{
	RefundStatusNoApply:       "没有申请退款",
	RefundStatusBuyerApplied:  "买家已经申请退款，等待卖家同意",
	RefundStatusSellerAgreed:  "卖家已经同意退款，等待买家退货",
	RefundStatusBuyerReturned: "买家已经退货，等待卖家确认收货",
	RefundStatusSuccess:       "退款成功",
	RefundStatusSellerRefused: "卖家拒绝退款",
	RefundStatusClosed:        "退款关闭",
}

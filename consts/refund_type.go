package consts

// RefundType represents refund types
const (
	RefundTypeOnlyRefund   = 1 // 仅退款 / Refunds only
	RefundTypeReturnRefund = 2 // 退货退款 / Returned items and refunds
)

// RefundTypeText maps refund type codes to their descriptions (Chinese + English)
var RefundTypeText = map[int]string{
	RefundTypeOnlyRefund:   "仅退款 / Refunds only",
	RefundTypeReturnRefund: "退货退款 / Returned items and refunds",
}

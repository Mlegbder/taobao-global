package consts

const (
	MessageTypeStatusChange = 3 //订单状态变更
	MessageTypePriceChange  = 8 //订单改价变更
	MessageTypeRefund       = 9 //订单逆向单/留言同步消息
)

// MessageTypeDesc 消息类型对应的中文描述
var MessageTypeDesc = map[int]string{
	MessageTypeStatusChange: "订单状态变更",
	MessageTypePriceChange:  "订单改价变更",
	MessageTypeRefund:       "订单逆向单/留言同步消息",
}

// MessageTypeList 消息类型枚举列表（方便校验或遍历）
var MessageTypeList = []int{
	MessageTypeStatusChange,
	MessageTypePriceChange,
	MessageTypeRefund,
}

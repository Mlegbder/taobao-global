package consts

// LogisticsStatus 物流状态/节点状态枚举
const (
	// 运输中
	LogisticsStatusDelivering = "DELIVERING"

	// 已签收
	LogisticsStatusSigned = "SIGNED"

	// 已揽收
	LogisticsStatusCollected = "COLLECTED"

	// 已拒收
	LogisticsStatusRejected = "REJECTED"

	// 已退回
	LogisticsStatusReturned = "RETURNED"

	// 异常
	LogisticsStatusException = "EXCEPTION"

	// 待揽收
	LogisticsStatusWaitingCollect = "WAITING_COLLECT"

	// 待清关
	LogisticsStatusWaitingCustoms = "WAITING_CUSTOMS"

	// 清关中
	LogisticsStatusCustomsClearance = "CUSTOMS_CLEARANCE"

	// 清关失败
	LogisticsStatusCustomsFailed = "CUSTOMS_FAILED"

	// 其他未知状态
	LogisticsStatusUnknown = "UNKNOWN"
)

// LogisticsStatusDesc 状态码对应的中文描述
var LogisticsStatusDesc = map[string]string{
	LogisticsStatusDelivering:       "运输中",
	LogisticsStatusSigned:           "已签收",
	LogisticsStatusCollected:        "已揽收",
	LogisticsStatusRejected:         "已拒收",
	LogisticsStatusReturned:         "已退回",
	LogisticsStatusException:        "异常",
	LogisticsStatusWaitingCollect:   "待揽收",
	LogisticsStatusWaitingCustoms:   "待清关",
	LogisticsStatusCustomsClearance: "清关中",
	LogisticsStatusCustomsFailed:    "清关失败",
	LogisticsStatusUnknown:          "未知状态",
}

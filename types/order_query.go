package types

// QueryPurchaseOrdersRequest 查询采购订单请求
type QueryPurchaseOrdersRequest struct {
	Status          string  `json:"status,omitempty"`            // 订单状态
	SortType        string  `json:"sort_type,omitempty"`         // ASC/DESC
	PageNo          int     `json:"page_no,omitempty"`           // 页码
	PageSize        int     `json:"page_size,omitempty"`         // 每页数量
	ModifyTimeStart int64   `json:"modify_time_start,omitempty"` // 开始修改时间 (UTC 时间戳)
	ModifyTimeEnd   int64   `json:"modify_time_end,omitempty"`   // 结束修改时间 (UTC 时间戳)
	OuterPurchaseID string  `json:"outer_purchase_id,omitempty"` // ISV 采购单ID
	PurchaseIDS     []int64 `json:"purchase_ids,omitempty"`      // 主单ID列表
}

// QueryPurchaseOrdersResponse 查询采购订单响应
type QueryPurchaseOrdersResponse struct {
	Success   bool                    `json:"success"`
	ErrorCode string                  `json:"error_code"`
	ErrorMsg  string                  `json:"error_msg"`
	Data      *PurchaseOrdersPageData `json:"data"`
}

// PurchaseOrdersPageData 订单分页结果
type PurchaseOrdersPageData struct {
	ResultsTotal   int64           `json:"results_total"`
	PageNo         int             `json:"page_no"`
	PageSize       int             `json:"page_size"`
	PurchaseOrders []PurchaseOrder `json:"purchase_orders"`
}

// PurchaseOrder 主单信息
type PurchaseOrder struct {
	Status            string             `json:"status"`
	OrderSource       string             `json:"order_source"`
	Receiver          *OrderAddress      `json:"receiver"`
	PayTime           string             `json:"pay_time"`
	PurchaseID        int64              `json:"purchase_id"`
	DistributorNick   string             `json:"distributor_nick"`
	PurchaseCurrency  string             `json:"purchase_currency"`
	OuterPurchaseID   string             `json:"outer_purchase_id"`
	PurchaseAmount    int64              `json:"purchase_amount"`
	SellerOrderNumber string             `json:"seller_order_number"`
	ModifyTime        int64              `json:"modify_time"`
	CreatedTime       int64              `json:"created_time"`
	PayCurrency       string             `json:"pay_currency"`
	ProductAmount     int64              `json:"product_amount"`
	SourceMarket      string             `json:"source_market"`
	SupplierNick      string             `json:"supplier_nick"`
	SubPurchaseOrders []SubPurchaseOrder `json:"sub_purchase_orders"`
}

// SubPurchaseOrder 子单信息
type SubPurchaseOrder struct {
	SubPurchaseOrderID int64  `json:"sub_purchase_order_id"`
	SkuID              int64  `json:"sku_id"`
	Status             string `json:"status"`
	Quantity           int    `json:"quantity"`
	Title              string `json:"title"`
	ItemID             string `json:"item_id"`
	PayAmount          string `json:"pay_amount"`
	PurchaseAmount     int64  `json:"purchase_amount"`
	PurchaseCurrency   string `json:"purchase_currency"`
	RefundID           int64  `json:"refund_id"`
	RefundStatus       int    `json:"refund_status"`
	ErrorCode          string `json:"error_code"`
	ErrorMessage       string `json:"error_message"`
}

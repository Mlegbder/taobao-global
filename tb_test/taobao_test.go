package tb_test

import (
	"fmt"
	"github.com/Mlegbder/taobao-global/taobao"
	"github.com/Mlegbder/taobao-global/types"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

// 内存存储实现
type MemoryTokenStore struct {
	token *types.TokenResponse
}

func (m *MemoryTokenStore) SaveToken(token *types.TokenResponse) error {
	m.token = token
	fmt.Println(token)
	return nil
}

func (m *MemoryTokenStore) LoadToken() (*types.TokenResponse, error) {
	if m.token == nil {
		return &types.TokenResponse{AccessToken: "dummy-token", RefreshToken: "dummy-refresh"}, nil
	}
	return m.token, nil
}

// 商品详情
func TestRunItemDetail(t *testing.T) {
	client := getClient()
	req := types.QueryAllProductRequest{
		ItemID: "755475833284",
	}
	resp, err := client.Item.GetDetail(req)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 商品详情
func TestRunItemSearch(t *testing.T) {
	client := getClient()
	req := types.ItemSearchRequest{
		Keyword:  "bags",
		PageNo:   1,
		PageSize: 10,
		Language: "en",
	}
	resp, err := client.Item.Search(req)
	if err != nil {
		log.Fatalf("❌ 商品搜索失败: %v", err)
	}
	fmt.Printf("✅ 搜索到 %d 条商品\n", len(resp.Data.Items))
}

// 获取货源详情(带翻译)
func TestRunSourceItemDetail(t *testing.T) {
	client := getClient()
	req := types.ItemDetailRequest{
		ItemResource: "taobao",
		ItemID:       "806339192392",
		Language:     "en",
	}
	resp, err := client.Item.GetSourceItemDetail(req)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 创建采购单
func TestRunCreateOrder(t *testing.T) {
	client := getClient()
	req := types.CreatePurchaseOrderRequest{
		OuterPurchaseID: "TESTMORE100000004",
		PurchaseAmount:  2000, // 单位: 分
		OrderLineList: []types.OrderLineReq{
			{ItemID: "4096709792297020",
				SkuID:       "32248575036476",
				Quantity:    1,
				Currency:    "CNY",
				Price:       1000,
				OrderLineNo: "TESTMORE100000004",
			},
			{ItemID: "4096612521451972",
				SkuID:       "30312388416964",
				Quantity:    1,
				Currency:    "CNY",
				Price:       1000,
				OrderLineNo: "TESTMORE100000004",
			},
		},
		Receiver: types.OrderAddress{
			Name:        "ProfessorWen",
			Country:     "中国大陆",
			State:       "广东省",
			City:        "广州市",
			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
			MobilePhone: "13068212342",
		},
		WarehouseAddressInfo: &types.OrderAddress{
			Name:        "ProfessorWen",
			Country:     "中国大陆",
			State:       "广东省",
			City:        "广州市",
			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
			MobilePhone: "13068212342",
		},
		OrderRemark: "Test order",
	}
	resp, err := client.Order.Create(req)
	if err != nil {
		log.Fatalf("❌ 创建采购单失败: %v", err)
	}
	if resp.Success {
		fmt.Printf("✅ 采购单创建成功: %s, 支付链接: %s\n",
			resp.Data.OuterPurchaseID, resp.Data.PaymentURL)
	} else {
		fmt.Printf("❌ 采购单创建失败: %s\n", resp.ErrorMsg)
	}
}

// 查询采购单
func TestRunQueryPurchaseOrders(t *testing.T) {
	client := getClient()
	req := types.QueryPurchaseOrdersRequest{
		PurchaseIDS: []int64{200077867837}, // 采购单ID
		PageNo:      1,
		PageSize:    10,
	}

	resp, err := client.Order.Query(req)
	if err != nil {
		log.Fatalf("❌ 查询采购单失败: %v", err)
	}

	if resp.Success {
		fmt.Printf("✅ 共查询到 %d 个采购单\n", resp.Data.ResultsTotal)
		for _, order := range resp.Data.PurchaseOrders {
			fmt.Printf("📦 主单ID: %d, 状态: %s, 金额: %.2f %s\n",
				order.PurchaseID,
				order.Status,
				float64(order.PurchaseAmount)/100,
				order.PurchaseCurrency,
			)
			for _, sub := range order.SubPurchaseOrders {
				fmt.Printf("   - 子单ID: %d, 商品: %s, 数量: %d, 状态: %s\n",
					sub.SubPurchaseOrderID, sub.Title, sub.Quantity, sub.Status)
			}
		}
	} else {
		fmt.Printf("❌ 查询失败: %s (%s)\n", resp.ErrorMsg, resp.ErrorCode)
	}
}

// 查询物流详情
func TestRunLogisticsDetail(t *testing.T) {
	client := getClient()
	req := types.GetLogisticsDetailRequest{
		PurchaseOrderLineID: 2912066688476065752,
	}
	resp, err := client.Logistics.GetDetail(req)
	if err != nil {
		log.Fatalf("❌ 查询采购单失败: %v", err)
	}
	fmt.Println(resp)
}

// 查询退款详情
func TestRunQueryRefundOrder(t *testing.T) {
	client := getClient()
	req := types.QueryRefundOrderRequest{
		RefundID: 110006829192,
	}
	resp, err := client.Order.QueryRefundOrder(req)
	if err != nil {
		log.Fatalf("❌ 查询退款单失败: %v", err)
	}
	fmt.Println(resp)
}

// 查询采购账单
func TestRunQueryPurchaseBill(t *testing.T) {
	client := getClient()
	req := types.PurchaseBillRequest{
		TimeType:  "paytime",
		PageNo:    1,
		PageSize:  10,
		StartTime: 1756698588,
		EndTime:   1757562605,
	}
	resp, err := client.Bill.PurchaseBill(req)
	if err != nil {
		log.Fatalf("❌ 查询采购账单失败: %v", err)
	}
	fmt.Println(resp)
}

// 查询退款账单
func TestRunQueryRefundBill(t *testing.T) {
	client := getClient()
	req := types.RefundBillRequest{
		PageNo:    "1",
		PageSize:  "10",
		StartTime: "1756698588",
		EndTime:   "1757562605",
	}
	resp, err := client.Bill.RefundBill(req)
	if err != nil {
		log.Fatalf("❌ 查询退款账单失败: %v", err)
	}
	fmt.Println(resp)
}

// 获取客户端
func getClient() *taobao.Client {
	// 1. 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: .env file not found, will use system environment variables")
	}

	// 2. 从环境变量读取
	appKey := os.Getenv("TAOBAO_APP_KEY")
	appSecret := os.Getenv("TAOBAO_APP_SECRET")
	accessToken := os.Getenv("TAOBAO_ACCESS_TOKEN")
	refreshToken := os.Getenv("TAOBAO_REFRESH_TOKEN")
	baseApi := os.Getenv("TAOBAO_BASE_API")

	store := &MemoryTokenStore{
		token: &types.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	client := taobao.NewClient(baseApi, appKey, appSecret, store)
	return client
}

package tb_test

import (
	"fmt"
	"github.com/Mlegbder/taobao-global/consts"
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
		ItemID: "731040994692",
	}
	resp, err := client.Item.GetDetail(req)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 商品搜索
func TestRunItemSearch(t *testing.T) {
	client := getClient()
	req := types.ItemSearchRequest{
		PageNo:   1,
		PageSize: 10,
		Language: "en",
		ShopID:   599014143,
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
		ItemID:       "731040994692",
		Language:     "en",
	}
	resp, err := client.Item.GetSourceItemDetail(req)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 批量支付
func TestRunBatchPay(t *testing.T) {
	client := getClient()
	req := types.BatchPayPurchaseOrderRequest{
		PurchaseOrderIDList: []int64{200078851363}, //采购IDS
	}

	resp, err := client.Order.BatchPay(req)
	if err != nil {
		log.Fatalf("batch pay failed: %v", err)
	}

	fmt.Println(resp)
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
		PurchaseIDS: []int64{200078283966}, // 采购单ID
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
		RefundID: 110006916474,
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
		TimeType:  consts.TimeTypeCreate,
		PageNo:    1,
		PageSize:  10,
		StartTime: 1756665600000,
		EndTime:   1757529600000,
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
		StartTime: "1756665600000",
		EndTime:   "1757529600000",
	}
	resp, err := client.Bill.RefundBill(req)
	if err != nil {
		log.Fatalf("❌ 查询退款账单失败: %v", err)
	}
	fmt.Println(resp)
}

// 图片上传
func TestRunImageUpload(t *testing.T) {
	client := getClient()
	// 假设你已经把图片转成 Base64 字符串
	imgBase64 := "iVBORw0KGgoAAAANSUhEUgAAA5QAAANbCAIAAACy6kd+AAAAC...."

	req := types.ImageUploadRequest{
		ImageBase64: imgBase64,
	}

	resp, err := client.Upload.Image(req)
	if err != nil {
		log.Fatalf("❌ 图片上传失败: %v", err)
	}

	if resp.Data != nil {
		fmt.Printf("✅ 图片上传成功, ImageID: %s\n", resp.Data.ImageID)
	} else {
		fmt.Printf("❌ 上传失败: %s (%s)\n", resp.BizErrorMsg, resp.BizErrorCode)
	}
}

// 图片搜索
func TestRunImgSearch(t *testing.T) {
	client := getClient()
	// 用 image_id 搜索 (推荐：先调用 ImageUpload 上传图片获取 image_id)
	req := types.ImgSearchRequest{
		PicURL:   "https://womata-gr.oss-accelerate.aliyuncs.com/lianfei/1756290364/TRANS-O1CN015sGm5922HZ8MoAfT9_2212513067095-0-cib-f.jpg",
		Language: "en",
	}

	resp, err := client.Item.ImgSearch(req)
	if err != nil {
		log.Fatalf("❌ 图片搜索失败: %v", err)
	}

	if len(resp.Data) > 0 {
		fmt.Printf("✅ 找到 %d 个商品\n", len(resp.Data))
		for _, item := range resp.Data {
			fmt.Printf("- %s (ID: %d, 价格: %s 元)\n", item.Title, item.ItemID, item.Price)
		}
	} else {
		fmt.Println("未找到相关商品")
	}
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

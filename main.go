package main

import (
	"fmt"
	"github.com/Mlegbder/taobao-global/taobao"
	"github.com/Mlegbder/taobao-global/types"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	BaseApi = "https://api.taobao.global/rest"
)

// main 函数选择要执行的示例
func main() {
	// 获取客户端
	client, accessToken := getClient()

	// 关键词查询商品
	// runItemSearch(client, accessToken)

	// 获取商品详情
	// runItemDetail(client, accessToken)

	// 订单预览
	// runOrderPreview(client, accessToken)

	// 创建采购单
	// runCreateOrder(client, accessToken)

	// 取消采购单
	//runCancelOrder(client, accessToken)

	// 批量支付
	// runBatchPay(client, accessToken)

	// 查询采购单物流详情
	// runGetLogisticsDetail(client, accessToken)

	// 查询采购单
	runQueryPurchaseOrders(client, accessToken)

	//图片上传
	// runImageUpload(client, accessToken)

	// 图片搜索
	//runImgSearch(client, accessToken)
}

// ========== 示例函数们 ==========

// 商品搜索
func runItemSearch(client *taobao.Client, accessToken string) {
	req := types.ItemSearchRequest{
		Keyword:  "bags",
		PageNo:   1,
		PageSize: 5,
		Language: "en",
	}
	resp, err := client.Item.Search(req, accessToken)
	if err != nil {
		log.Fatalf("❌ 商品搜索失败: %v", err)
	}
	fmt.Printf("✅ 搜索到 %d 条商品\n", len(resp.Data.Items))
}

// 商品详情
func runItemDetail(client *taobao.Client, accessToken string) {
	req := types.ItemDetailRequest{
		ItemResource: "taobao",
		ItemID:       "956575770963",
		Language:     "en",
	}
	resp, err := client.Item.GetDetail(req, accessToken)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 订单预览
func runOrderPreview(client *taobao.Client, accessToken string) {
	req := types.PurchaseOrderRenderRequest{
		NeedSupplyChainService: false,
		RenderItemList:         `[{"item_id":"4096526553499286","sku_id":"28464810350230","quantity":2}]`,
		WarehouseAddress: &types.Address{
			Name:        "ProfessorWen",
			Country:     "中国大陆",
			State:       "广东省",
			City:        "广州市",
			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
			MobilePhone: "13068212342",
		},
		ReceiverAddress: types.Address{
			Name:        "ProfessorWen",
			Country:     "中国大陆",
			State:       "广东省",
			City:        "广州市",
			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
			MobilePhone: "13068212342",
		},
	}
	resp, err := client.Order.Render(req, accessToken)
	if err != nil {
		log.Fatalf("❌ 订单预览失败: %v", err)
	}
	fmt.Printf("✅ 订单预览成功: %+v\n", resp.Data)
}

// 创建采购单
func runCreateOrder(client *taobao.Client, accessToken string) {
	req := types.CreatePurchaseOrderRequest{
		OuterPurchaseID: "ISV123456789",
		PurchaseAmount:  199600, // 单位: 分
		OrderLineList: `[{
			"item_id": "4096526553499286", 
			"sku_id": "28464810350230", 
			"quantity": 2,
			"currency": "CNY",
			"price": 199600,
			"order_line_no": "ISV123456789"
		}]`,
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
	resp, err := client.Order.Create(req, accessToken)
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

// 批量支付
func runBatchPay(client *taobao.Client, accessToken string) {
	req := types.BatchPayPurchaseOrderRequest{
		PurchaseOrderIDList: []int64{202509020001, 202509020002}, //采购IDS
	}

	resp, err := client.Order.BatchPay(req, accessToken)
	if err != nil {
		log.Fatalf("batch pay failed: %v", err)
	}

	if resp.Success {
		fmt.Println("✅ 批量支付任务已提交")
		fmt.Println("待支付订单: ", resp.Data.WillPayPurchaseOrderIDs)
		if len(resp.Data.PayFailurePurchaseOrderIDs) > 0 {
			fmt.Println("❌ 支付失败订单: ", resp.Data.PayFailurePurchaseOrderIDs)
		}
	} else {
		fmt.Printf("❌ 批量支付失败: %s (%s)\n", resp.ErrorMsg, resp.ErrorCode)
	}
}

// 取消采购单
func runCancelOrder(client *taobao.Client, accessToken string) {
	req := types.AsynCancelPurchaseOrderRequest{
		PurchaseID:   "200077684761", // 替换成真实采购单号
		CancelReason: "Customer requested cancellation",
		CancelRemark: "测试取消订单",
	}
	resp, err := client.Order.AsynCancel(req, accessToken)
	if err != nil {
		log.Fatalf("❌ 取消订单失败: %v", err)
	}
	if resp.Success {
		fmt.Println("✅ 取消订单请求已发起 (异步)，请调用 /purchase/orders/query 查询最终状态")
	} else {
		fmt.Printf("❌ 取消订单失败: %s (%s)\n", resp.ErrorMsg, resp.ErrorCode)
	}
}

// 查询采购单物流详情
func runGetLogisticsDetail(client *taobao.Client, accessToken string) {
	req := types.GetLogisticsDetailRequest{
		PurchaseOrderLineID: 1234567890, // 子单号
	}

	resp, err := client.Logistics.GetDetail(req, accessToken)
	if err != nil {
		log.Fatalf("get logistics detail failed: %v", err)
	}

	if resp.Success {
		fmt.Printf("✅ 当前物流状态: %s (%s)\n", resp.Data.LogisticsDesc, resp.Data.LogisticsStatus)
		for _, pkg := range resp.Data.PnmLogisticsDetails {
			fmt.Printf("📦 包裹单号: %s\n", pkg.MailNo)
			for _, trace := range pkg.LogisticsTraces {
				fmt.Printf("   [%s] %s - %s (%s)\n",
					trace.Time, trace.Status, trace.StatusDesc, trace.City)
			}
		}
	} else {
		fmt.Printf("❌ 查询失败: %s (%s)\n", resp.ErrorMsg, resp.ErrorCode)
	}
}

// 查询采购单
func runQueryPurchaseOrders(client *taobao.Client, accessToken string) {
	req := types.QueryPurchaseOrdersRequest{
		PurchaseIDS: []int64{200077684761}, // 采购单ID
		PageNo:      1,
		PageSize:    10,
	}

	resp, err := client.Order.Query(req, accessToken)
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

// 图片上传
func runImageUpload(client *taobao.Client, accessToken string) {
	// 假设你已经把图片转成 Base64 字符串
	imgBase64 := "UklGRpxpAgBXRUJQVlA4WAoAAAAgAAAArwQArwQASUND..."

	req := types.ImageUploadRequest{
		ImageBase64: imgBase64,
	}

	resp, err := client.Upload.Image(req, accessToken)
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
func runImgSearch(client *taobao.Client, accessToken string) {
	// 用 image_id 搜索 (推荐：先调用 ImageUpload 上传图片获取 image_id)
	req := types.ImgSearchRequest{
		ImageID:  "1521908561144519126",
		Language: "en",
	}

	resp, err := client.Item.ImgSearch(req, accessToken)
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

// ========== 工具函数 ==========

// 获取客户端
func getClient() (*taobao.Client, string) {
	// 1. 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: .env file not found, will use system environment variables")
	}

	// 2. 从环境变量读取
	appKey := os.Getenv("TAOBAO_APP_KEY")
	appSecret := os.Getenv("TAOBAO_APP_SECRET")
	accessToken := os.Getenv("TAOBAO_ACCESS_TOKEN")
	if appKey == "" || appSecret == "" || accessToken == "" {
		log.Fatal("❌ TAOBAO_APP_KEY / TAOBAO_APP_SECRET / TAOBAO_ACCESS_TOKEN is not set")
	}

	client := taobao.NewClient(BaseApi, appKey, appSecret)
	return client, accessToken
}

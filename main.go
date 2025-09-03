package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"taobao-global/taobao"
	"taobao-global/types"
)

// 商品搜索示例
//func main() {
//	//获取客户端
//	client, accessToken := getClient()
//
//	req := types.ItemSearchRequest{
//		Keyword:  "bags",
//		PageNo:   1,
//		PageSize: 10,
//		Language: "en",
//	}
//
//	resp, err := client.Item.Search(req, accessToken)
//	if err != nil {
//		log.Fatalf("get detail failed: %v", err)
//	}
//
//	fmt.Println("商品: ", resp.Data)
//}

// 商品详情示例
//func main() {
//	//获取客户端
//	client, accessToken := getClient()
//
//	req := types.ItemDetailRequest{
//		ItemResource: "taobao",
//		ItemID:       "956575770963",
//		Language:     "en",
//	}
//
//	resp, err := client.Item.GetDetail(req, accessToken)
//	if err != nil {
//		log.Fatalf("get detail failed: %v", err)
//	}
//
//	fmt.Printf("商品标题: %s\n", resp.Data.Title)
//	fmt.Printf("优惠价: %.2f 元\n", float64(resp.Data.PromotionPrice)/100)
//}

// 订单预览示例 item_id 和 sku_id就是详情返回的mpId和mpSkuId
//func main() {
//	//获取客户端
//	client, accessToken := getClient()
//	req := types.PurchaseOrderRenderRequest{
//		NeedSupplyChainService: false,
//		RenderItemList:         `[{"item_id":"4096526553499286","sku_id":"28464810350230","quantity":2}]`,
//		WarehouseAddress: &types.Address{
//			Name:        "ProfessorWen",
//			Country:     "中国大陆",
//			State:       "广东省",
//			City:        "广州市",
//			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
//			MobilePhone: "13068212342",
//		},
//		ReceiverAddress: types.Address{
//			Name:        "ProfessorWen",
//			Country:     "中国大陆",
//			State:       "广东省",
//			City:        "广州市",
//			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
//			MobilePhone: "13068212342",
//		},
//	}
//
//	resp, err := client.Order.Render(req, accessToken)
//	if err != nil {
//		log.Fatalf("order preview failed: %v", err)
//	}
//
//	fmt.Println("预览订单: ", resp.Data)
//}

// 创建采购单示例
func main() {
	//获取客户端
	client, accessToken := getClient()
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
		log.Fatalf("create order failed: %v", err)
	}

	if resp.Success {
		fmt.Printf("采购单创建成功, 采购ID: %s, 支付链接: %s\n",
			resp.Data.OuterPurchaseID, resp.Data.PaymentURL)
	} else {
		fmt.Printf("采购单创建失败: %s\n", resp.ErrorMsg)
	}
}

func getClient() (*taobao.Client, string) {
	// 1. 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, will use system environment variables")
	}

	// 2. 从环境变量读取
	appKey := os.Getenv("TAOBAO_APP_KEY")
	appSecret := os.Getenv("TAOBAO_APP_SECRET")
	accessToken := os.Getenv("TAOBAO_ACCESS_TOKEN")
	if appKey == "" || appSecret == "" || accessToken == "" {
		log.Fatal("❌ TAOBAO_APP_KEY or TAOBAO_APP_SECRET or TAOBAO_ACCESS_TOKEN is not set")
	}
	client := taobao.NewClient(
		appKey,
		appSecret,
	)
	return client, accessToken
}

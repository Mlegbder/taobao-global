package main

import (
	"fmt"
	"github.com/Mlegbder/taobao-global/taobao"
	"github.com/Mlegbder/taobao-global/types"
	"log"
)

const (
	BaseApi = "https://api.taobao.global/rest"
)

// main 函数选择要执行的示例
func main() {

	// 关键词查询商品
	// runItemSearch(client)

	// 获取商品详情
	// runItemDetail(client)

	// 获取商品翻译
	// runItemTranslate(client)

	// 订单预览
	// runOrderPreview(client)

	// 创建采购单
	// runCreateOrder(client)

	// 取消采购单
	// runCancelOrder(client)

	// 批量支付
	// runBatchPay(client)

	// 查询采购单物流详情
	// runGetLogisticsDetail(client)

	// 查询采购单
	// runQueryPurchaseOrders(client)

	//图片上传
	// runImageUpload(client)

	// 图片搜索
	//runImgSearch(client)
}

// ========== 示例函数们 ==========

// 商品搜索
func runItemSearch(client *taobao.Client) {
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

// 商品详情
func runItemDetail(client *taobao.Client) {
	req := types.QueryAllProductRequest{
		ItemID: "805577403719",
	}
	resp, err := client.Item.GetDetail(req)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 商品货源详情
func runSourceItemDetail(client *taobao.Client) {
	req := types.ItemDetailRequest{
		ItemResource: "taobao",
		ItemID:       "778127375879",
		Language:     "en",
	}
	resp, err := client.Item.GetSourceItemDetail(req)
	if err != nil {
		log.Fatalf("❌ 商品详情获取失败: %v", err)
	}
	fmt.Printf("✅ 商品标题: %s, 优惠价: %.2f 元\n",
		resp.Data.Title, float64(resp.Data.PromotionPrice)/100)
}

// 商品翻译
func runItemTranslate(client *taobao.Client) {
	req := types.ProductTranslateRequest{
		ItemID:   "4096623585210707", // mp_id
		Language: "en",
	}

	resp, err := client.Item.Translate(req)
	if err != nil {
		log.Fatalf("❌ 商品翻译失败: %v", err)
	}

	fmt.Println(resp)
}

// 订单预览
func runOrderPreview(client *taobao.Client) {
	req := types.PurchaseOrderRenderRequest{
		NeedSupplyChainService: false,
		RenderItemList: []types.RenderItemReq{
			{ItemID: "4096526553499286", SkuID: "28464810350230", Quantity: 2},
		},
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
			City:        "广州市",
			Address:     "白云湖街道机场路兵房街兵工厂67号集运仓",
			MobilePhone: "13068212342",
		},
	}
	resp, err := client.Order.Render(req)
	if err != nil {
		log.Fatalf("❌ 订单预览失败: %v", err)
	}
	fmt.Printf("✅ 订单预览成功: %+v\n", resp.Result)
}

// 创建采购单
func runCreateOrder(client *taobao.Client) {
	req := types.CreatePurchaseOrderRequest{
		OuterPurchaseID: "TEST100000001",
		PurchaseAmount:  2000, // 单位: 分
		OrderLineList: []types.OrderLineReq{
			{ItemID: "4096701167701319",
				SkuID:       "32077491877191",
				Quantity:    1,
				Currency:    "CNY",
				Price:       1000,
				OrderLineNo: "TEST100000001",
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

// 批量支付
func runBatchPay(client *taobao.Client) {
	req := types.BatchPayPurchaseOrderRequest{
		PurchaseOrderIDList: []int64{200077867837}, //采购IDS
	}

	resp, err := client.Order.BatchPay(req)
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
func runCancelOrder(client *taobao.Client) {
	req := types.AsynCancelPurchaseOrderRequest{
		PurchaseID:   "200077821489", // 替换成真实采购单号
		CancelReason: "Customer requested cancellation",
		CancelRemark: "测试取消订单",
	}
	resp, err := client.Order.AsynCancel(req)
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
func runGetLogisticsDetail(client *taobao.Client) {
	req := types.GetLogisticsDetailRequest{
		PurchaseOrderLineID: 1234567890, // 子单号
	}

	resp, err := client.Logistics.GetDetail(req)
	if err != nil {
		log.Fatalf("get logistics detail failed: %v", err)
	}
	fmt.Println(resp)
}

// 查询采购单
func runQueryPurchaseOrders(client *taobao.Client) {
	req := types.QueryPurchaseOrdersRequest{
		PurchaseIDS: []int64{200077684761}, // 采购单ID
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

// 图片上传
func runImageUpload(client *taobao.Client) {
	// 假设你已经把图片转成 Base64 字符串
	imgBase64 := "UklGRpxpAgBXRUJQVlA4WAoAAAAgAAAArwQArwQASUND..."

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
func runImgSearch(client *taobao.Client) {
	// 用 image_id 搜索 (推荐：先调用 ImageUpload 上传图片获取 image_id)
	req := types.ImgSearchRequest{
		ImageID:  "1521908561144519126",
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

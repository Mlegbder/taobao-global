## 🚀 功能特性

- ✅ `TokenService.Create` 获取 `access_token`
- ✅ `ItemService.Search` 商品搜索
- ✅ `...` 更多接口引入SDK查看
- ✅ 参数自动签名（HMAC-SHA256）
- ✅ 强类型 Request/Response，IDE 自动补全
- ✅ 支持手动传入 `access_token`，方便多账号/缓存管理

---

## 📦 安装

在项目中执行：

```bash
go get github.com/Mlegbder/taobao-global
````

---

## 🛠 使用示例

```go
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
   client := getClient()

   // 关键词查询商品
   runItemSearch(client)

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

   if resp.Success && resp.Data != nil {
      fmt.Printf("✅ 商品标题 (%s): %s\n", resp.Data.Language, resp.Data.Title)
      for _, prop := range resp.Data.Properties {
         fmt.Printf(" - %s: %s\n", prop.PropName, prop.ValueName)
      }
   } else {
      fmt.Printf("❌ 翻译失败: %s (%s)\n", resp.ErrorMsg, resp.ErrorCode)
   }

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
   fmt.Printf("✅ 订单预览成功: %+v\n", resp.Data)
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

// 查询退款单
func runQueryRefundOrder(client *taobao.Client) {
   req := types.QueryRefundOrderRequest{
      RefundID: 1234567890,
   }

   resp, err := client.Order.QueryRefundOrder(req)
   if err != nil {
      log.Fatalf("❌ 查询退款单失败: %v", err)
   }

   if resp.Success && resp.Data != nil {
      fmt.Printf("✅ 退款单 %d 状态: %d\n", resp.Data.RefundOrder.RefundID, resp.Data.RefundOrder.RefundStatus)
      fmt.Printf("退款金额: %.2f 元\n", float64(resp.Data.RefundOrder.RefundFee)/100)
      if resp.Data.PurchaseOrderLine != nil {
         fmt.Printf("商品: %s, 数量: %d\n",
            resp.Data.PurchaseOrderLine.ItemTitle,
            resp.Data.PurchaseOrderLine.Quantity,
         )
      }
   } else {
      fmt.Printf("❌ 查询失败: %s (%s)\n", resp.ErrorMsg, resp.ErrorCode)
   }
}

// ========== 工具函数 ==========

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
   if appKey == "" || appSecret == "" || accessToken == "" {
      log.Fatal("❌ TAOBAO_APP_KEY / TAOBAO_APP_SECRET / TAOBAO_ACCESS_TOKEN is not set")
   }

   client := taobao.NewClient(BaseApi, appKey, appSecret, accessToken)
   return client
}



```

---

## 📂 项目结构

```
├── README.md              # 项目说明文档
├── go.mod                 # Go 模块配置
├── main.go                # 使用示例
├── consts/                # 常量定义（API 地址等）
│   └── api.go
├── taobao/                # SDK 核心封装
│   ├── client.go          # Client 主入口
│   ├── item_service.go    # 商品搜索服务
│   └── token_service.go   # Token 管理服务
│   └── order_service.go   # 订单服务
│   └── ....go   # 其他请求服务
├── types/                 # 请求/响应数据类型
│   ├── base.go
│   ├── item_search.go
│   └── token.go
│   └── ....go   # 其他请求数据类型
└── utils/                 # 工具方法
    └── taobao_util.go     # 签名 & HTTP 请求
```

---

## ⚠️ 注意事项

1. **access\_token 不要硬编码**

    * 通过 `client.Token.Create` 获取后，请保存到数据库或缓存。
    * 业务请求时，调用 `client.Item.Search(req)` 并传入缓存的 token。

2. **Token 生命周期**

    * `access_token` 有效期为 **30 天**, `refresh_token` 有效期为 **60 天**。
    * 过期后可用 `refresh_token` 获取新的 `access_token`。

3. **请求签名**

    * 所有请求参数会自动进行 HMAC-SHA256 签名。

4. **多账号支持**

    * SDK 不会缓存 token，你可以在同一个 `Client` 下传入不同的 `access_token`，以支持多店铺。

---

## 📄 License

MIT License

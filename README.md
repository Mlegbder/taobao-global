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

### 1. 获取 Access Token

```go
package main

import (
    "taobao-global/types"
    "fmt"
    "log"
    "os"
)

func main() {
    client := taobao.NewClient(
        os.Getenv("TAOBAO_APP_KEY"),
        os.Getenv("TAOBAO_APP_SECRET"),
    )

    tokenResp, err := client.Token.Create(types.TokenRequest{
        Code: "your_oauth_code_here",
    })
    if err != nil {
        log.Fatalf("token request failed: %v", err)
    }

    // 保存到数据库 / Redis
    saveTokenToDB(tokenResp.AccessToken, tokenResp.RefreshToken)

    fmt.Println("AccessToken:", tokenResp.AccessToken)
}
```

---

### 2. 搜索商品

```go
package main

import (
    "taobao-global/types"
    "fmt"
    "log"
)

func main() {
    client := taobao.NewClient(
        "your_app_key",
        "your_app_secret",
    )

    // 从数据库 / Redis 加载 token
    accessToken := loadTokenFromDB()

    req := types.ItemSearchRequest{
        Keyword:  "运动鞋",
        Sort:     "SALE_QTY_DESC",
        PageNo:   1,
        PageSize: 10,
        Filters:  []string{"min_price:5000", "max_price:20000"}, // 单位: 分
        Language: "en",
    }

    resp, err := client.Item.Search(req, accessToken)
    if err != nil {
        log.Fatalf("item search failed: %v", err)
    }

    fmt.Printf("Found %d items\n", len(resp.Data.Items))
    for _, item := range resp.Data.Items {
        fmt.Printf("[%d] %s - %s元 (店铺: %s)\n", item.ItemID, item.Title, item.Price, item.ShopName)
    }
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
    * 业务请求时，调用 `client.Item.Search(req, accessToken)` 并传入缓存的 token。

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

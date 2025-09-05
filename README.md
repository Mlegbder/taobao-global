## 🚀 功能特性


- ✅ 参数自动签名（HMAC-SHA256）
- ✅ 强类型 Request/Response，IDE 自动补全

---

## 📦 安装

在项目中执行：

```bash
go get github.com/Mlegbder/taobao-global
````

---

## 🛠 使用示例
自定义实现SaveToken和LoadToken, SDK内置自动刷新access_token

```go
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

func TestClientExecute(t *testing.T) {
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


## 📄 License

MIT License

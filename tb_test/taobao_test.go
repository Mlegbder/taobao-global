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
		ItemID: "805577403719",
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

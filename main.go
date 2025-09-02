package main

import (
	"fmt"
	"log"
	"os"
	"taobao-global/taobao"
	"taobao-global/types"
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

	fmt.Println("AccessToken:", tokenResp.AccessToken)
}

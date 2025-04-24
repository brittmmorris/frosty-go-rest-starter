package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	routerID := os.Getenv("ROUTER_ID")
	routerKey := os.Getenv("ROUTER_KEY")
	prompt := "Tell me a joke"
	rule := "cost"

	if routerID == "" || routerKey == "" {
		fmt.Println("Missing ROUTER_ID or ROUTER_KEY")
		return
	}

	baseURL := "https://api.gofrosty.ai/chat"
	params := url.Values{}
	params.Add("router_id", routerID)
	params.Add("router_key", routerKey)
	params.Add("prompt", prompt)
	params.Add("rule", rule)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}

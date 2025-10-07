package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// MicroCMSのAPIレスポンスの構造体
type TopPageContent struct {
	// Goのinterface{} (任意の型) を使って柔軟にすべてのフィールドを受け取れるようにします。
	Contents map[string]interface{} `json:",inline"`
}

func main() {
	// 環境変数の確認
	serviceID := os.Getenv("MICROCMS_SERVICE_ID")
	apiKey := os.Getenv("MICROCMS_API_KEY")

	if serviceID == "" || apiKey == "" {
		fmt.Println("Error: MICROCMS_SERVICE_ID or MICROCMS_API_KEY not set.")
		os.Exit(1)
	}

	// MicroCMS APIのURLを構築
	apiURL := fmt.Sprintf("https://%s.microcms.io/api/v1/top_page", serviceID)

	// HTTPリクエストの作成
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	// API Keyをリクエストヘッダーに設定
	req.Header.Set("X-MICROCMS-API-KEY", apiKey)

	// APIの実行
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API returned status code %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// データを読み込み
	body, _ := ioutil.ReadAll(resp.Body)
	
	// 保存先のディレクトリを作成
	dataDir := filepath.Join(".", "data", "cms")
	os.MkdirAll(dataDir, 0755)

	// JSONファイルをローカルに保存
	savePath := filepath.Join(dataDir, "top_page.json")
	if err := ioutil.WriteFile(savePath, body, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fetched and saved data to %s\n", savePath)
}
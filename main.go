package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	return port
}

type Symbol struct {
	BaseCoin            string   `json:"baseCoin"`
	BuyLimitPriceRatio  string   `json:"buyLimitPriceRatio"`
	FeeRateUpRatio      string   `json:"feeRateUpRatio"`
	LimitOpenTime       string   `json:"limitOpenTime"`
	MaintainTime        string   `json:"maintainTime"`
	MakerFeeRate        string   `json:"makerFeeRate"`
	MinTradeNum         string   `json:"minTradeNum"`
	OffTime             string   `json:"offTime"`
	OpenCostUpRatio     string   `json:"openCostUpRatio"`
	PriceEndStep        string   `json:"priceEndStep"`
	PricePlace          string   `json:"pricePlace"`
	QuoteCoin           string   `json:"quoteCoin"`
	SellLimitPriceRatio string   `json:"sellLimitPriceRatio"`
	SizeMultiplier      string   `json:"sizeMultiplier"`
	SupportMarginCoins  []string `json:"supportMarginCoins"`
	Symbol              string   `json:"symbol"`
	SymbolName          string   `json:"symbolName"`
	SymbolStatus        string   `json:"symbolStatus"`
	SymbolType          string   `json:"symbolType"`
	TakerFeeRate        string   `json:"takerFeeRate"`
	VolumePlace         string   `json:"volumePlace"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "HELLO WORLD PLEASE WORK",
		})
	})

	url := "https://api.bitget.com/api/mix/v1/market/contracts?productType=umcbl"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending GET request: %s", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error reading response body: %s", err.Error())
		return
	}

	type Data struct {
		Data []map[string]interface{} `json:"data"`
	}

	var data Data
	err2 := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err2)
		return
	}

	for _, item := range data.Data {
		baseCoin := item["baseCoin"].(string)
		// Extract other fields as needed
		fmt.Println("Base Coin:", baseCoin)
	}
	// Access the "data" key using type assertion
	// dataValue, ok := data["data"].(map[string]interface{})
	// if !ok {
	// 	fmt.Println("Failed to assert data type")
	// 	fmt.Println(data)

	// 	return
	// }

	// Print the value of the "data" key
	// fmt.Println(dataValue)

	app.Listen("0.0.0.0" + getPort())

}

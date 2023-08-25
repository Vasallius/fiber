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

func processBitcoinData(data [][]interface{}) {
	closePrices := make([]string, 0)

	for _, item := range data {
		// Access individual elements of each item in the data array
		timestamp := item[0]
		open := item[1]
		high := item[2]
		low := item[3]
		close := item[4].(string)

		// Do something with the data, such as printing it
		fmt.Println("Timestamp:", timestamp)
		fmt.Println("Open:", open)
		fmt.Println("High:", high)
		fmt.Println("Low:", low)
		fmt.Println("Close:", close)
		closePrices = append(closePrices, close)
		// Add your own logic here to process the data, err.Error()
	}
	fmt.Println(closePrices)
}

func getSymbols() error {
	url := "https://fapi.binance.com/fapi/v1/exchangeInfo"
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error sending GET request: %w", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var exchangeInfo map[string]interface{}
	if err := json.Unmarshal(body, &exchangeInfo); err != nil {
		return fmt.Errorf("error parsing JSON response: %w", err)
	}

	// Do something with the symbols...

	return nil
}

func getBitcoinData() error {
	url := "https://fapi.binance.com/fapi/v1/klines?symbol=BTCUSDT&interval=5m&limit=20"

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error sending GET request: %s", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %s", err.Error())
	}
	var data [][]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		println("error parsing JSON response", err.Error())
		return fmt.Errorf("error parsing JSON response: %s", err.Error())
	}

	// Call the processBitcoinData function with the retrieved data
	processBitcoinData(data)

	return nil
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "HELLO WORLD PLEASE WORK",
		})
	})
	getBitcoinData()
	app.Listen("0.0.0.0" + getPort())

}

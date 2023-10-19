package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CoinInfo struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
	LastUpdated  string  `json:"last_updated"`
}

func fetchCoinData(coinSymbol string) (*CoinInfo, error) {
	c := &http.Client{}
	var data []CoinInfo
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	for _, coin := range data {
		if coin.Symbol == coinSymbol {
			return &coin, nil
		}
	}

	return nil, fmt.Errorf("Криптовалюта c символом %s не найдена", coinSymbol)
}

func main() {
	var coinSymbol string
	fmt.Scan(&coinSymbol)
	for {
		coinData, err := fetchCoinData(coinSymbol)
		if err != nil {
			fmt.Println("Ошибка при получении данных:", err)
		} else {
			fmt.Printf("Курс криптовалюты %s (%s): $%.4f USD \nПоследнее обновление: %s\n", coinData.Name, coinData.Symbol, coinData.CurrentPrice, coinData.LastUpdated)
		}

		time.Sleep(10 * time.Minute)
	}
}

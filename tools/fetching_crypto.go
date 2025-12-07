package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FetchCryptoDataInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var FetchCryptoInfo = FetchCryptoDataInfo{
	Name:        "FetchCryptoData",
	Description: "Fetches the current price of a specified cryptocurrency in a given fiat currency.",
}

type FetchCryptoDataParams struct {
	Coin     string `json:"crypto_coin"`
	Currency string `json:"currency"`
}

type FetchCryptoDataReturn struct {
	Price float64 `json:"price"`
	Error string  `json:"error,omitempty"`
}

func FetchCryptoData(coin string, currency string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", coin, currency)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	price, ok := result[coin][currency]
	if !ok {
		return 0, fmt.Errorf("price data not found for %s in %s", coin, currency)
	}

	return price, nil
}

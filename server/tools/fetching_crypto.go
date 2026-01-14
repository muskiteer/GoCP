package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
)


func FetchCryptoData(ctx context.Context, args map[string]any) (any, error) {
	coin, _ := args["coin"].(string)
	currency, _ := args["currency"].(string)
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
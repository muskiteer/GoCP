package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"strings"

)


func FetchCryptoData(ctx context.Context, args map[string]any) (any, error) {
	
	coin, ok := args["crypto_name"].(string)
	if !ok || coin == "" {
		return nil, fmt.Errorf("crypto_name must be a non-empty string")
	}

	currency, ok := args["currency"].(string)
	if !ok || currency == "" {
		return nil, fmt.Errorf("currency must be a non-empty string")
	}

	currency = strings.ToLower(currency)
	coin = strings.ToLower(coin)
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", coin, currency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("coingecko API returned status %d", resp.StatusCode)
	}

	var result map[string]map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	rawprice, ok := result[coin][currency]
	
	if !ok {
		return nil, fmt.Errorf("price data not found for %s in %s", coin, currency)
	}


	price := fmt.Sprintf("%v", rawprice)

	return price, nil
}
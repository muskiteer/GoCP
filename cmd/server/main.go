package main

import (
	"github.com/muskiteer/GoCP/tools"
	"log"
)

func main(){
	price, err := tools.FetchCryptoData("bitcoin", "inr")
	if err != nil {
		log.Fatalf("Error fetching crypto data: %v", err)
	}
	log.Printf("The current price of Bitcoin in USD is: %f", price)
}
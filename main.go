package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/rafaelandrade/API-RedCoins/api"
)

func main() {
	api.Run()

	coin := os.Args[1]
	coins := strings.Split(coin, ",")
	fmt.Println("")
	fmt.Printf("[CoinMarketCap] %s\n", coins)
	for _, c := range coins {
		result := getCoins(c)
		fmt.Println("---------------------------")
		fmt.Println(`Price(BTC):  `, result.PriceBtc)
		fmt.Println(`Price(USD):  `, result.PriceUsd)
		fmt.Println(`Volume:      `, result.Two4HVolumeUsd)
		fmt.Println(`Change 24H:  `, result.PercentChange24H)
		fmt.Println("")
	}
}
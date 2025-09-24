package main

import (
	"fmt"
	"log"

	"github.com/handruka/arbitrage-bot/internal/config"
	fetcher "github.com/handruka/arbitrage-bot/internal/fetcher"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка LoadConfig файла:", err)
		fmt.Println("Ошибка загрузки .env файла")
	}

	fmt.Println(cfg)

	f := fetcher.NewFetcher(cfg.BinanceApiKey, cfg.BinanceApiSecret)

	price, err := f.GetPrices(cfg.Pairs)
	if err != nil {
		log.Fatal("Ошибка: %v", err)
	}

	fmt.Println(price)
}

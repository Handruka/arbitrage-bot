package main

import (
	"fmt"
	"log"

	"github.com/handruka/arbitrage-bot/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка LoadConfig файла:", err)
		fmt.Println("Ошибка загрузки .env файла")
	}

	fmt.Println(cfg)
}

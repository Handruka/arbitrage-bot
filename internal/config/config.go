package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BinanceApiKey    string
	BinanceApiSecret string
	Pairs            string
	StartAmount      float64
	Fee              float64
}

func NewConfig() Config {
	return Config{}
}

func LoadConfig() (Config, error) {

	cfg := NewConfig()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
		return cfg, err
	}

	cfg.BinanceApiKey = os.Getenv("BINANCE_API_KEY")
	cfg.BinanceApiSecret = os.Getenv("BINANCE_API_SECRET")
	cfg.Pairs = os.Getenv("PAIRS")
	cfg.StartAmount, err = strconv.ParseFloat(os.Getenv("START_AMOUNT"), 64)
	if err != nil {
		log.Fatal("Ошибка чтения ключа StartAmount .env файла:", err)
		return cfg, err
	}
	cfg.Fee, err = strconv.ParseFloat(os.Getenv("FEE"), 64)
	if err != nil {
		log.Fatal("Ошибка чтения ключа Fee .env файла:", err)
		return cfg, err
	}

	return cfg, nil

}

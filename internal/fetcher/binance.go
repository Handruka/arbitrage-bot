package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Fetcher struct {
	ApiKey    string
	ApiSecret string
	BaseURL   string
	Client    *http.Client
}

type PriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func NewFetcher(apiKey, apiSecret string) *Fetcher {
	return &Fetcher{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		BaseURL:   "https://api.binance.com",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (f *Fetcher) GetPrices(pairs []string) (map[string]float64, error) {

	price := make(map[string]float64)

	for _, v := range pairs {

		var builder strings.Builder

		builder.WriteString(f.BaseURL)
		builder.WriteString("/api/v3/ticker/price?symbol=")
		builder.WriteString(v)
		url := builder.String()

		resp, err := f.Client.Get(url)
		if err != nil {
			log.Printf("Не удаётся сделать запрос к binance, ошибка: %v", err)
			return price, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			return price, fmt.Errorf("binance http %v: %s", resp.StatusCode, string(b))
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Не удаётся прочитать тело ответа сервера binance, ошибка: %v", err)
			return price, err
		}

		var priceObj PriceResponse

		err = json.Unmarshal(body, &priceObj)
		if err != nil {
			log.Printf("Не удаётся декодировать json-ответ сервера binance, ошибка: %v", err)
			return price, err
		}
		priceFlt, err := strconv.ParseFloat(priceObj.Price, 64)
		if err != nil {
			log.Printf("Не удаётся конвертировать json-ответ сервера binance Price в float64, ошибка: %v", err)
			return price, err
		}
		price[priceObj.Symbol] = priceFlt

	}
	return price, nil
}

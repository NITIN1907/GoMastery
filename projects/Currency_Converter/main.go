package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type exchangeRate struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func convertCurrency(from string, to string, amount float64) {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/f3c91ee3d299bfc1efe4cd93/latest/%s", from)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()

	var data exchangeRate
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("error decoding:", err)
		return
	}

	rate, ok := data.ConversionRates[to]
	if !ok {
		fmt.Println("Currency is not supported...")
		return
	}

	result := amount * rate
	fmt.Printf("%.2f %s = %.2f %s (rate: %.4f)\n", amount, from, result, to, rate)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <from_currency> <to_currency> <amount>")
		return
	}
	from := os.Args[1]
	to := os.Args[2]
	var amount float64
	fmt.Sscanf(os.Args[3], "%f", &amount)
	convertCurrency(from, to, amount)
}

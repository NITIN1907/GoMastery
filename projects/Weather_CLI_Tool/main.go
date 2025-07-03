package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getWeather(city string, apiKey string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error..")
		return
	}
	defer res.Body.Close()

	var data Weather
	json.NewDecoder(res.Body).Decode(&data)
	fmt.Printf("Weather in %s: %.2f°C\n", data.Name, data.Main.Temp)
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city>")
		return
	}
	apiKey := "abc123yourkeyhere456"
	city := os.Args[1]
	getWeather(city, apiKey)
}

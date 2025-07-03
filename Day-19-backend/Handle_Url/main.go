package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myurl := "https://example.com:8080/search?q=golang&sort=asc"

	parseURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("format %T\n", parseURL)
	fmt.Println("Schema: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Query: ", parseURL.RawQuery)
	fmt.Println("Port: ", parseURL.Port())
	fmt.Println("Host name: ", parseURL.Hostname())

	urls := "https://httpbin.org/get"
	res, _ := http.Get(urls)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

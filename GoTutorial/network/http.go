package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello Golang!")

	response, err := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=BTC")

	if err != nil {
		fmt.Printf("Request faild with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

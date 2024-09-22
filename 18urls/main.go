package main

import (
	"fmt"
	"net/url"
)

const URLS = "https://example.com:8080/search?q=chatgpt&lang=en"

func main() {
	welcome := "Welcome to the url handing in golang"
	fmt.Println(welcome)
	fmt.Println(URLS)
	result, err := url.Parse(URLS)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Path)
	fmt.Println(result.RawPath)
	// fmt.Println(result.Port())
	// qparams := result.Query()
	// for _, val := range qparams {
	// 	fmt.Println("The parameters are:", val)
	// }
	partofurls := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/search",
	}
	fmt.Println(partofurls)

}

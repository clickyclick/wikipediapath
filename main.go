package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/net/html"
)

func main() {
	url := "https://en.wikipedia.org/wiki/Go_(programming_language)"
    resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
}
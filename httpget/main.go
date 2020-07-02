package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func httpGet() {
	resp, err := http.Get("http://csharp.news")
	check(err)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Printf("%s\n", string(data))
}

func httpClient() {
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	check(err)

	req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}

func main() {
	httpGet()
	httpClient()
}

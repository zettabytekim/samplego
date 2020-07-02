package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

func httpPostForm() {
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	check(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func httpPost() {
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	check(err)

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	httpPost()
	httpPostForm()
}

package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

func xmlPost() {
	person := Person{"Alex", 10}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	check(err)
	req.Header.Add("Context-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
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
	xmlPost()
}

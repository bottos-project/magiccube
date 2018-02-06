package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://www.yinyuetai.com/mv/get-bigpic"

func main()  {
	post()
}

func get()  {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func post()  {
	resp, err := http.Post(url, "application/json",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

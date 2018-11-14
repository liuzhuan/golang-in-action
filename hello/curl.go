package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	url := "https://baidu.com/"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
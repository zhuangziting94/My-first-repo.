package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://muxithief.muxixyz.com/api/v1/login"
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

//{"data":"","msg":"警告!缺少code"}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8080/videos"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
	fmt.Println(string(body))

}

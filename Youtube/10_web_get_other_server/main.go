package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Disease_id   int
	Disease_name string
	Count        int `json:"COUNT(name)"`
}

type Response struct {
	Success bool
	Message string
	Count   int
	Data    []Data
}

func main() {
	resp, _ := http.Get("http://localhost:3001/api/v1/topTenDisease")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var Response Response
	json.Unmarshal(bytes, &Response)       //byte to JSON
	fmt.Printf("Response : %+v", Response) //Json formate

	resp.Body.Close()
}

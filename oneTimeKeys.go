package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"bytes"
)

func main() {
	url := "https://sandbox-api-pay.line.me/v2/payments/oneTimeKeys/pay"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"productName": "test product","amount": 100,"currency": "TWD","orderId": "TW2020-LINE-00004","oneTimeKey": "389253804433039534"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LINE-ChannelId", "1649580251")
	req.Header.Set("X-LINE-ChannelSecret", "ddca54d0f3e50847af3f6ec4fcaef890")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
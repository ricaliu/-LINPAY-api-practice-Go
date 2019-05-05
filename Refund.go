package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"bytes"
)

func main() {
	url := "https://sandbox-api-pay.line.me/v2/payments/orders/TW2019-LINE-00003/refund"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"refundAmount": 35}`)
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
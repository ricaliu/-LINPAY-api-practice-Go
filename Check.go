package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	url := "https://sandbox-api-pay.line.me/v2/payments/orders/TW2019-LINE-00003/check"
    fmt.Println("URL:>", url)

  
    req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LINE-ChannelId", "1649580251")
	req.Header.Set("X-LINE-ChannelSecret", "ddca54d0f3e50847af3f6ec4fcaef890")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close() //The client must close the response body when finished with it

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
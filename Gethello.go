package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	url := "http://rica.topedu.io:5000/v2/payments/oneTimeKeys"
    fmt.Println("URL:>", url)

  
    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close() //The client must close the response body when finished with it

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
package main

import (
    "fmt"
    "net/http"
    "time"
    //"encoding/json"
    "io"
)

func main() {
    http.HandleFunc("/test", proxy)
    http.ListenAndServe(":8019", nil)
}

func proxy(w http.ResponseWriter, req *http.Request) {
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", "https://postman-echo.com/headers", nil)
    req.Header.Add("key","the-secret-key")
    resp,err:=client.Do(req)
    output, err := io.ReadAll(resp.Body)
    fmt.Fprintf(w, "%s",output)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}

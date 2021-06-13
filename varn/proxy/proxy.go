package main

import (
    "fmt"
    "net/http"
    "time"
    "io"
    "flag"
)

func main() {
    var host="localhost"
    var port=8080
    flag.IntVar(&port, "port", 8080, "specify port to use, defaults to 8080.")
    flag.Parse()
    http.HandleFunc("/test", proxy)
    addr := fmt.Sprintf("%s:%d", host, port)
    http.ListenAndServe(addr, nil)
}

func proxy(w http.ResponseWriter, req *http.Request) {
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", "https://postman-echo.com/headers", nil)
    if err != nil {
        panic(err)
    }
    req.Header.Add("key","the-secret-key")
    resp,err:=client.Do(req)
    if err != nil {
        panic(err)
    }
    output, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, "%s",output)
    defer resp.Body.Close()
}

package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
)

func main() {
    http.HandleFunc("/test", proxy)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":8017", nil)
}

var result interface {}

func proxy(w http.ResponseWriter, req *http.Request) {
    var result interface {}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, . Welcome!")
    fmt.Fprintf(w, "%s",message)

    //resp, err := http.Get("https://postman-echo.com/headers")
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", "https://postman-echo.com/headers", nil)
    req.Header.Add("key","the-secret-key")
    resp,err:=client.Do(req)
    decoder:=json.NewDecoder(resp.Body)
    if err:=decoder.Decode(&result); err!=nil {
         panic (err)
    }
    //n=result
    fmt.Fprintf(w, "%s",result)
    //body, err := io.ReadAll(resp.Body)

    fmt.Fprintf(w, "%s",resp)
    //fmt.Fprintf(w, "%s",body['headers'])

    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

package main
import (
    "flag"
    "fmt"
    "time"
    "net/http"
    "io"
)
func main() {
    var host="http://localhost"
    var port=8080
    flag.IntVar(&port, "port", 8080, "specify port to use, defaults to 8080.")
    flag.Parse()
    addr := fmt.Sprintf("%s:%d/test", host, port)

    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest("GET", addr, nil)
    if err != nil {
        panic(err)
    }
    resp,err:=client.Do(req)
    if err != nil {
        panic(err)
    }
    output, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(output))
    defer resp.Body.Close()

}

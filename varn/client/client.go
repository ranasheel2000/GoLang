package main

import (
    "fmt"

    "github.com/ranasheel2000/GoLang/varn/proxy"
)

func main() {
    // Get a greeting message and print it.
    message := proxy.Hello("Gladys")
    fmt.Println(message)
}


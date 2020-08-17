package main

import (
	"fmt"
	"os"
)

func main() {
	var args string
	for _, y := range os.Args[1:] {
		args += " " + y
		fmt.Print(y)
	}

}

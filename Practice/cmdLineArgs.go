package main

import (
	"fmt"
	"os"
)

func main() {
	var allArgs string = ""
	for i := 1; i < len(os.Args); i++ {
		allArgs += "     " + os.Args[i]
	}

	fmt.Println("allArgs=", allArgs)
}

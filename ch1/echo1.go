package main

import (
	"fmt"
	"os"
)

func main() {
	var sep, o string
	for i := 1; i < len(os.Args); i++ {
		o += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(o)
}

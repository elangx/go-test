package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	counts := make(map[string]int)
	for _, v := range os.Args[1:] {
		f, err := os.Open(v)
		if err != nil {
			fmt.Println("error")
			continue
		}
		countLines(f, counts)
	}

	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

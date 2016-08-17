package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, v := range os.Args[1:] {
		file, err := ioutil.ReadFile(v)
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(file), "\n") {
			if len(line) > 0 {
				counts[line]++
			}
		}
	}

	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

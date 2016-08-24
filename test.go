package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type tPoint struct {
	Point
	s int
}

func main() {
	c := tPoint{}
	fmt.Println(c.X)
}

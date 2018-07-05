package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 7}
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())
}

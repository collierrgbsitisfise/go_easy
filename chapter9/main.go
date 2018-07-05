package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	a := Circle{0, 0, 7}
	b := Circle{1, 2, 5}
	c := Circle{1, 3, 10}
	fmt.Println(totalArea(&c, &b, &a))
	fmt.Println(c.area())
}

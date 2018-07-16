package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUB
)

func main() {
	var a [3]int = [3]int{3, 4, 1}
	q := [...]int{4, 5, 6}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	fmt.Println(q)

	symbol := [...]string{USD: "$", EUR: "E", GBP: "R", RUB: "R"}
	fmt.Println(symbol[RUB])
}

package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {
	x := 5
	zero(&x)
	y := new(int)
	zero(y)
	fmt.Println(x)
	fmt.Println(*y)
}

package main

import "fmt"

func swap(x *int, y *int) {
	tmp := new(int)
	*tmp = *x
	*x = *y
	*y = *tmp
}

func main() {
	x := 5
	y := 10
	fmt.Println("Before swaping")
	fmt.Println("x ", x)
	fmt.Println("y ", y)
	swap(&x, &y)
	fmt.Println("x ", x)
	fmt.Println("y ", y)

}

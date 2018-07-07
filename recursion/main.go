package main

import "fmt"

func countDown(num int) int {
	if num == 0 {
		return num
	}

	fmt.Println(num)
	return countDown(num - 1)
}

func main() {
	fmt.Println(countDown(10))
}

package main

import (
	"fmt"
)

func sumOfArr(arr []int, index int) int {
	if index == len(arr)-1 {
		return arr[len(arr)-1]
	}

	return arr[index] + sumOfArr(arr, index+1)
}

func main() {
	arr := []int{1, 3, 4, 6}
	fmt.Println(sumOfArr(arr, 0))
}

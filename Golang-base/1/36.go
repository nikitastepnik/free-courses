package main

import "fmt"

func main() {
	var A int
	var workArray []int
	fmt.Scan(&A)
	for A > 0 {
		workArray = append(workArray, A%2)
		A /= 2
	}
	for i := len(workArray) - 1; i >= 0; i-- {
		fmt.Print(workArray[i])
	}
}

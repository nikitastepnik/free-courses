package main

import "fmt"

func main() {
	array := []int{}
	var n int
	var c int
	var countPositiveNumbers int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		array = append(array, c)
		if array[i] > 0 {
			countPositiveNumbers += 1
		}
	}
	fmt.Println(countPositiveNumbers)

}

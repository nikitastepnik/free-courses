package main

import "fmt"

func main() {
	array := []int{}
	var n int
	var c int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		array = append(array, c)
		if i%2 == 0 {
			fmt.Print(array[i], " ")
		}
	}

}

package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	var max int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
		if i == 0 {
			max = array[i]
		}
		if a > max {
			max = a
		}
	}
	fmt.Println(max)

}

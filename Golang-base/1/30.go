package main

import "fmt"

func main() {
	var N int
	var number int
	var min int
	var count int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&number)
		if i == 0 {
			min = number
			count += 1
		} else {
			if min > number {
				min = number
				count = 1
			} else if number == min {
				count += 1
			}
		}
	}
	fmt.Println(count)

}

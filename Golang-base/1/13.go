package main

import "fmt"

func main() {
	var n int
	var max int
	var count int

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			count = 0
		}
		if n == max {
			count += 1
		}
	}
	fmt.Println(count)
}

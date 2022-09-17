package main

import "fmt"

func main() {
	var N int
	var number int
	var counts = 0
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&number)
		if number == 0 {
			counts += 1
		}
	}
	fmt.Println(counts)

}

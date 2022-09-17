package main

import "fmt"

func main() {
	var n, countDegrees int
	var initialNubm = 1
	fmt.Scan(&n)

	for i := n; i > 0; {
		i /= 2
		countDegrees += 1
	}
	for i := 0; i < countDegrees; i++ {
		fmt.Print(initialNubm, " ")
		initialNubm *= 2
	}
}

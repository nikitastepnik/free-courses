package main

import "fmt"

func minimumFromFour() int {
	var min, curNumb int
	fmt.Scan(&min)

	for i := 0; i < 3; i++ {
		fmt.Scan(&curNumb)
		if curNumb < min {
			min = curNumb
		}
	}

	return min
}





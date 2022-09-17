package main

import "fmt"

func main() {
	var a, b, maxDerivSeven int

	fmt.Scan(&a)
	fmt.Scan(&b)
	maxDerivSeven = a
	for i := a; i <= b; i++ {
		if i%7 == 0 && i > maxDerivSeven {
			maxDerivSeven = i
		}
	}
	if maxDerivSeven == a && maxDerivSeven%7 != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(maxDerivSeven)

	}
}

package main

import "fmt"

func main() {
	var n int

OuterLoop:
	for {
		fmt.Scan(&n)
		if n < 10 {
			continue
		}
		if n > 100 {
			break OuterLoop
		} else {
			fmt.Println(n)
		}
	}

}

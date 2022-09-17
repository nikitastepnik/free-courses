package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	if n == 1 {
		fmt.Println("1 korova")
	} else if n < 5 {
		fmt.Printf("%d korovy", n)
	} else if n < 20 {
		fmt.Printf("%d korov", n)
	} else {
		switch n % 10 {
		case 1:
			fmt.Printf("%d korova", n)
		case 2:
			fmt.Printf("%d korovy", n)
		case 3:
			fmt.Printf("%d korovy", n)
		case 4:
			fmt.Printf("%d korovy", n)
		default:
			fmt.Printf("%d korov", n)
		}
	}
}

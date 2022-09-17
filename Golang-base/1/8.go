package main

import "fmt"

func main() {

	var number int
	fmt.Scan(&number)

	if (number%10 != number/10%10) && (number%10 != number/100) && (number/10%10 != number/100) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

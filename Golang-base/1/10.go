package main

import "fmt"

func main() {

	var number int
	fmt.Scan(&number)

	firstHalfOfNumber := number / 1000
	secondHalfOfNumber := number % 1000
	sumNumbersOfFirstHalf := firstHalfOfNumber/100 + firstHalfOfNumber%10 + firstHalfOfNumber/10%10
	sumNumbersOfSecondHalf := secondHalfOfNumber/100 + secondHalfOfNumber%10 + secondHalfOfNumber/10%10
	if sumNumbersOfFirstHalf == sumNumbersOfSecondHalf {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

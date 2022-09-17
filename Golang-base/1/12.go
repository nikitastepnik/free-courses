package main

import "fmt"

func main() {
	var countNumbers int
	var curNumber int
	var sumNumbersMultiplesEight int
	var i int
	fmt.Scan(&countNumbers)

	for i < countNumbers {
		fmt.Scan(&curNumber)
		if curNumber%8 == 0 && curNumber < 100 && curNumber > 9 {
			sumNumbersMultiplesEight += curNumber
		}
		i++
	}
	fmt.Println(sumNumbersMultiplesEight)
}

package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	firstNumber := number % 10
	secondNumber := number / 10 % 10
	thirdNumber := number / 100
	fmt.Printf("%d%d%d", firstNumber, secondNumber, thirdNumber)

}

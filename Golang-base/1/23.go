package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	sumNumbers := number%10 + number/10%10 + number/100
	fmt.Println(sumNumbers)

}

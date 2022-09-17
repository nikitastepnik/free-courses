package main

import "fmt"

func main() {

	var a int
	var b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли

	b = (a % 100) / 10

	fmt.Println(b)
}

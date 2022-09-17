package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}

}

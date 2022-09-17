package main

import "fmt"

func main() {
	var countElements int
	fmt.Scan(&countElements)
	var a []int
	var b int
	for i := 0; i < countElements; i++ {
		fmt.Scan(&b)
		a = append(a, b)
	}
	fmt.Println(a[3])
}

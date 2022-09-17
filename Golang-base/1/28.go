package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	//c := int((a + b) / 2)
	//d := float32((a + b) / 2)
	if (a+b)%2 == 0 {
		fmt.Println((a + b) / 2)
	} else {
		fmt.Println(float32(a+b) / 2)
	}

}

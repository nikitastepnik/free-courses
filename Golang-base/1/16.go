package main

import (
	"fmt"
)

func main() {
	var initialSum float32
	var percent float32
	var finalSum float32
	var countYear int
	fmt.Scan(&initialSum)
	fmt.Scan(&percent)
	fmt.Scan(&finalSum)
	var percentil = 1 + percent/100.0
	for i := initialSum; i <= finalSum; {
		i = float32(int(i * percentil))
		countYear += 1
	}
	fmt.Println(countYear)

}

package main

import (
	"fmt"
)

func countRyazryads(b int) (razryads int) {
	if b == 10000 {
		razryads = 5
	} else if b/1000 != 0 {
		razryads = 4
	} else if b/100 != 0 {
		razryads = 3
	} else if b/10 != 0 {
		razryads = 2
	} else {
		razryads = 1
	}
	return razryads
}

func evaluateMin(a, b int) (min int) {
	if a > b {
		return b
	} else {
		return a
	}

}

func main() {
	var a int
	var b int
	var razryadsA int
	var razryadsB int

	fmt.Scan(&a)
	fmt.Scan(&b)
	razryadsA = countRyazryads(a)
	razryadsB = countRyazryads(b)
	minRazryad := evaluateMin(razryadsA, razryadsB)
	s := make(
		[]int,
		minRazryad,
		minRazryad,
	)
	for i := 1; i <= razryadsA; i++ {
		curDigit := a % 10
		copyB := b
		for c := 1; c <= razryadsB; c++ {
			if curDigit == copyB%10 {
				s = append(s, curDigit)
				break
			}
			copyB /= 10
		}
		a /= 10
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 {
			fmt.Print(s[i], " ")
		}
	}
}

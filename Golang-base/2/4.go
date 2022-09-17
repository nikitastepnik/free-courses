package main

func sumInt(a ...int) (int, int) {
	var sum int
	var counts int
	for _, elem := range a {
		sum += elem
		counts += 1
	}
	return counts, sum
}

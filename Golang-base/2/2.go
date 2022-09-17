package main

func vote(x int, y int, z int) int {
	var sum = x + y + z
	if sum > 1 {
		return 1
	} else {
		return 0
	}
}

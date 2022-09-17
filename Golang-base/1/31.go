package main

import "fmt"

func solveTask(N int) int {
	var workArray []int
	var sum int
	for N > 0 {
		workArray = append(workArray, N%10)
		N /= 10
	}
	for i := len(workArray) - 1; i > -1; i-- {
		sum += workArray[i]
	}
	if sum/10 == 0 {
		return sum % 10
	} else {
		sum = solveTask(sum)
	}
	return sum % 10
}

func main() {
	var N, answer int
	fmt.Scan(&N)

	answer = solveTask(N)
	fmt.Println(answer)
}

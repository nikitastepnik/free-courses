package main

import "fmt"

func main() {
	var A int
	var workArray []int
	fmt.Scan(&A)

	workArray = append(workArray, 0, 1)

	var i = 1
	var j = 1
	for i < A {
		workArray = append(workArray, workArray[j]+workArray[j-1])
		i = workArray[len(workArray)-1]
		j++
	}

	if workArray[len(workArray)-1] == A {
		fmt.Println(len(workArray) - 1)
	} else {
		fmt.Println(-1)
	}
}

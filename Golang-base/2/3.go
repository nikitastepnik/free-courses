package main

func fibonacci(n int) int {
	var workArray []int
	workArray = append(workArray, 1, 1)
	var i = 1
	for i < n {
		if workArray[i] == n {
			break
		}
		workArray = append(workArray, workArray[i]+workArray[i-1])
		i += 1
	}
	return workArray[n-1]
}

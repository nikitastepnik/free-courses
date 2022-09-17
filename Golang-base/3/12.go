package main

func task(chanel chan int, n int) {
	chanel <- n + 1
}

package main

func main() {
	chanel := make(chan int)

	go func(chanel chan int) {
		//work() some work
		close(chanel)
	}(chanel)

	<-chanel

}

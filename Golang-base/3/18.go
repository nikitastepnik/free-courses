package main

func calculator2(arguments <-chan int, done <-chan struct{}) <-chan int {
	chanel := make(chan int)
	var elem, sum int

	go func(chan int) {
		defer close(chanel)
		for {
			select {
			case elem = <-arguments:
				sum += elem
			case <-done:
				chanel <- sum
				return
			}
		}
	}(chanel)

	return chanel
}

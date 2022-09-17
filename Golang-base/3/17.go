package main

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	chanelOutput := make(chan int)

	var elem int
	go func() {
		defer close(chanelOutput)
		select {
		case elem = <-firstChan:
			chanelOutput <- elem * elem
			break
		case elem = <-secondChan:
			chanelOutput <- elem * 3
			break
		case <-stopChan:
			break
		}
	}()

	return chanelOutput
}

package main

import "sync"

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	workMap1, workMap2 := make(map[int]int), make(map[int]int)
	wg, wg1 := new(sync.WaitGroup), new(sync.WaitGroup)
	mn1 := new(sync.Mutex)
	go func() {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(id int, mu *sync.Mutex) {
				defer wg.Done()
				mu.Lock()
				workMap1[id] = <-in1
				workMap2[id] = <-in2
				mu.Unlock()
				wg1.Add(1)
				go func(id1 int) {
					defer wg1.Done()
					workMap1[id1] = fn(workMap1[id1])
					workMap2[id1] = fn(workMap2[id1])
				}(id)

			}(i, mn1)
		}

		go func() {
			wg.Wait()

			for i := 0; i < n; i++ {
				out <- workMap1[i] + workMap2[i]
			}
		}()

	}()

}

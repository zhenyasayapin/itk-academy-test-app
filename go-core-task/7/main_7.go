package main

func mergeChannels(out chan int, done chan struct{}, channels ...chan int) {
	for _, ch := range channels {
		go mergeChannel(out, done, ch)
	}
}

func mergeChannel(out chan int, done chan struct{}, ch chan int) {
	for {
		select {
		case n := <-ch:
			out <- n
		case <-done:
			return
		}
	}
}

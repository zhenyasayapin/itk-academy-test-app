package main

import "testing"

func TestMerge(t *testing.T) {
	done := make(chan struct{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	mergedCh := make(chan int)

	go mergeChannels(mergedCh, done, ch1, ch2, ch3)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	<-mergedCh
	<-mergedCh
	<-mergedCh

	close(done)
}

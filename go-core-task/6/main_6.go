package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	ack := make(chan struct{})
	done := make(chan struct{})

	go random(ch, ack, done)

	nums := run(ch, ack, 3)

	close(done)

	fmt.Println(nums)
}

func random(ch chan int, ack chan struct{}, done chan struct{}) {
	for {
		select {
		case <-ack:
			ch <- rand.Intn(1000)
		case <-done:
			return
		}
	}
}

func run(ch chan int, ack chan struct{}, count int) []int {
	var nums []int
	for range count {
		ack <- struct{}{}
		nums = append(nums, <-ch)
	}

	return nums
}

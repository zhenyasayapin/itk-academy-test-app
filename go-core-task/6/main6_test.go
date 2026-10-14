package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	ch := make(chan int)
	ack := make(chan struct{})
	done := make(chan struct{})

	go random(ch, ack, done)

	var nums []int
	nums = run(ch, ack, 3)
	close(done)

	if len(nums) != 3 {
		t.Errorf("Expecting length 3, got %d", len(nums))
	}
}

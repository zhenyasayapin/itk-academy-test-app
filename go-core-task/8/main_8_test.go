package main

import (
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	swg := New(3)

	for range 6 {
		swg.Add()
		go func(swg *SemaphoreWaitGroup) {
			defer swg.Done()
			time.Sleep(2 * time.Second)
		}(swg)
	}

	swg.Wait()
}

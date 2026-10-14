package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func conveyour(uintCh chan uint8, floatCh chan float64) {
	for {
		value, ok := <-uintCh
		if !ok {
			return
		}

		floatCh <- math.Pow(float64(value), 3)
	}
}

func main() {
	uintCh := make(chan uint8)
	floatCh := make(chan float64)

	defer close(uintCh)
	defer close(floatCh)

	go conveyour(uintCh, floatCh)

	for range 10 {
		uintCh <- uint8(rand.Intn(10))
		fmt.Println(<-floatCh)
		time.Sleep(time.Second)
	}
}

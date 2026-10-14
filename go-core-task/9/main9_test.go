package main

import "testing"

func TestConveyor(t *testing.T) {
	tests := []struct {
		value    uint8
		newValue float64
	}{
		{uint8(2), float64(8)},
		{uint8(0), float64(0)},
	}

	uintCh := make(chan uint8)
	floatCh := make(chan float64)

	defer close(uintCh)
	defer close(floatCh)

	go conveyour(uintCh, floatCh)

	for _, tt := range tests {
		uintCh <- tt.value
		newValue := <-floatCh
		if tt.newValue != newValue {
			t.Errorf("want %f, got %f", tt.newValue, newValue)
		}
	}
}

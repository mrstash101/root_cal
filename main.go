package main

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func FastInverseSqrt(x float32) float32 {
	const threehalfs = float32(1.5)

	x2 := x * float32(0.5)
	y := x
	i := *(*int32)(unsafe.Pointer(&y))
	i = 0x5f3759df - i>>1
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threehalfs - (x2 * y * y))
	return y
}

func TestFastInverseSqrt(t *testing.T) {
	tests := []struct {
		in, out float32
		e       float32
	}{
		{2.0, 1.0/1.414, 1e-3},
		{3.0, 1.0/1.732, 1e-3},
		{5.0, 1.0/2.236, 1e-3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("1/√%.2f", tt.in), func(t *testing.T) {
			result := FastInverseSqrt(tt.in)
			e := float32(math.Abs(float64(result - tt.out)))
			if tt.e < e {
				t.Errorf("want %f, but got %f", tt.out, result)
			}
		})

	}
}

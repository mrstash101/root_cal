package main

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func FastInverseSqrt(x float32) float32 {
	x2 := x * float32(0.5)
	 := x

	 i := *(*int32)(unsafe.Pointer(&y))
	 i = 0x5f3759df - i>>1
	 y = *(*float32)(unsafe.Pointer(&i))

	 y = y * (float32(1.5) - (x2 * y * y))
	 return y 
}

func FastInverseSqrt2(x float32) float32 {
	return float32(1.0 / math.Sqrt(float64(x)))
}



func TestFastInverseSqrt(t *testing.T) {
	datas := []struct{
		in float32
	}{
	{2.0},
	{3.0},
	{4.0},
	{5.0},
}
for _, tt := range datas {
	t.Run(fmt.Sprintf("1/√%.2f", tt.in), func(t *testing.T){
		result1 := FastInverseSqrt(tt.in)
		result2 := FastInverseSqrt2(tt.in)

		e := math.Abs(float64(result1 - result2))
		if 1e-3 < e  {
			t.Errorf("something wrong")
		}
	})
}
}


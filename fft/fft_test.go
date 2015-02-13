package fft

import (
	"math"
	"math/cmplx"
	"math/rand"
	"testing"
)

func normalFourier(data []complex128) {
	tmp := make([]complex128, len(data))
	for k := 0; k < len(data); k++ {
		tmp[k] = 0
		for i := 0; i < len(data); i++ {
			tmp[k] += data[i] *
				cmplx.Exp(complex(0, 2*float64(i*k)*math.Pi/float64(len(data))))
		}
	}
	for i, v := range tmp {
		data[i] = v
	}
}

func TestFft(t *testing.T) {
	data1 := make([]complex128, 32)
	data2 := make([]complex128, 32)
	for i := range data1 {
		r := complex(rand.Float64()*2-1, rand.Float64()*2-1)
		data1[i] = r
		data2[i] = r
	}
	Fft(data1)
	normalFourier(data2)
	for i := range data1 {
		diff := cmplx.Abs(data2[i] - data1[i])
		if diff > 0.001 {
			t.Errorf("Resulted value differs from normal fourir transform expected=%v, got=%v", data2[i], data1[i])
		}
	}
}

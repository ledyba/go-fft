package fft

import (
	"math"
	"math/cmplx"
)

func Fft(data []complex128) {
	if len(data) == 1 {
		return
	}
	hl := len(data) / 2
	even := make([]complex128, hl)
	odd := make([]complex128, hl)
	for i := 0; i < hl; i++ {
		even[i] = data[i] + data[i+hl]
		odd[i] = (data[i] - data[i+hl]) *
			cmplx.Exp(complex(0, 2*float64(i)*math.Pi/float64(len(data))))
	}
	Fft(even)
	Fft(odd)
	for i := 0; i < hl; i++ {
		data[2*i] = even[i]
		data[2*i+1] = odd[i]
	}
}

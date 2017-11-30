package dct

import "math"

// TODO: Implement faster version like FFT.

func DCT2D(dat []byte, w, h int) []float64 {
	freq := make([]float64, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			freq[y*w+x] = float64(dat[y*w+x])
		}
	}
	tmp := make([]float64, w)
	for y := 0; y < h; y++ {
		DCT1D(freq, w*y, 1, tmp, 0, 1, w)
		for x := 0; x < w; x++ {
			freq[w*y+x] = tmp[x]
		}
	}
	tmp = make([]float64, h)
	for x := 0; x < w; x++ {
		DCT1D(freq, x, w, tmp, 0, 1, h)
		for y := 0; y < h; y++ {
			freq[x*w+y] = tmp[y]
		}
	}
	return freq
}

func DCT1D(in []float64, iFrom, iStride int, out []float64, oFrom, oStride, length int) {
	n := length
	for k := 0; k < n; k++ {
		out[k] = 0
		for i := 0; i < n; i++ {
			out[oFrom+k*oStride] += float64(in[iFrom+(i*iStride)]) * dct(i, k, n)
		}
	}
}

func IDCT1D(in []float64, iFrom, iStride int, out []byte, oFrom, oStride, length int) {
	n := length
	for i := 0; i < n; i++ {
		v := 0.0
		for k := 0; k < n; k++ {
			v += float64(in[iFrom+(k*iStride)]) * dct(i, k, n)
		}
		out[oFrom+(i*oStride)] = byte(v)
	}
}

func dct(i, k, N int) float64 {
	c := math.Sqrt(1.0 / float64(N))
	if i != 0 {
		c *= 2
	}
	v := math.Cos(float64((2*i+1)*k) * math.Pi / float64(2*N))
	return c * v
}

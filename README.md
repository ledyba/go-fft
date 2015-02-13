# FFT implementation in Golang
Readable FFT implementation, but may not run fast...

## How to use
```golang
import "github.com/ledyba/go-fft/fft"


  data := make([]complex128, 32)
  for i := range data {
    // Fill data
    data[i] = complex(float64(i*2)/float64(32),0)
  }
  fft.Fft(data)

```

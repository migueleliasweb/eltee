package benchmarks

import (
	"testing"
)

//BenchmarkCustomMinMax
func BenchmarkCustomMinMax(b *testing.B) {
	var data []float64

	for n := 0; n < b.N; n++ {
		data = append(data, float64(n))
	}

	max := data[0]
	min := data[0]

	b.StartTimer()
	for _, n := range data {
		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}
	b.StopTimer()

}

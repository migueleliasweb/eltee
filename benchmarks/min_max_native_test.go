package benchmarks

import (
	"math"
	"sort"
	"testing"
)

func BenchmarkMathMax(b *testing.B) {
	var data []float64

	for n := 0; n < b.N; n++ {
		data = append(data, float64(n))
	}

	max := data[0]
	min := data[0]
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, n := range data {
			max = math.Max(max, n)
			min = math.Min(min, n)
		}
	}
	b.StopTimer()

}

func BenchmarkSortMinMax5000(b *testing.B) {
	var data []float64

	for n := 0; n < b.N; n++ {
		data = append(data, float64(n))
	}

	b.StartTimer()

	sort.Float64s(data)
	// max := data[len(data)-1]
	// min := data[0]
	b.StopTimer()

}

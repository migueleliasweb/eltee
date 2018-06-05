package benchmarks

import (
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/stat"
)

//BenchmarkIntMapAppend Benchmarks append on maps
func BenchmarkGoNumStdDev5000(b *testing.B) {
	//slice on purpose
	var data []float64

	for n := 0; n < 5000; n++ {
		data = append(data, float64(n))
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stat.StdDev(data, nil)
	}
	b.StopTimer()

}

//BenchmarkIntMapAppend Benchmarks append on maps
func BenchmarkGoNumMean5000(b *testing.B) {
	//slice on purpose
	var data []float64

	for n := 0; n < 5000; n++ {
		data = append(data, float64(n))
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stat.Mean(data, nil)
	}
	b.StopTimer()

}

//BenchmarkGoNumMin5000 Benchmarks math.Min
func BenchmarkGoNumMin5000(b *testing.B) {
	var data []float64

	for n := range rand.Perm(5000) {
		data = append(data, float64(n))
	}

	b.StartTimer()
	max := float64(0)
	for n := 0; n < b.N; n++ {
		for _, n := range data {
			max = math.Min(max, n)
		}
	}
	b.StopTimer()

}

//BenchmarkIntMapAppend Benchmarks math.Max
func BenchmarkGoNumMax5000(b *testing.B) {
	var data []float64

	for n := range rand.Perm(5000) {
		data = append(data, float64(n))
	}

	b.StartTimer()
	max := float64(0)
	for n := 0; n < b.N; n++ {
		for _, n := range data {
			max = math.Max(max, n)
		}
	}
	b.StopTimer()

}

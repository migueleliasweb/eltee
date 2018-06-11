package benchmarks

import (
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

package benchmarks

import (
	"sync"
	"testing"
)

//BenchmarkIntSliceAppend Benchmarks append on slices
func BenchmarkIntSliceAppend(b *testing.B) {
	slice := make([]int, 0)

	for n := 0; n < b.N; n++ {
		slice = append(slice, 123)
	}

}

//BenchmarkIntMapAppend Benchmarks append on maps
func BenchmarkIntMapAppend(b *testing.B) {
	mmap := make(map[int]int)

	for n := 0; n < b.N; n++ {
		mmap[n] = n
	}

}

//BenchmarkIntMapAppend Benchmarks append on maps
func BenchmarkIntSyncMapAppend(b *testing.B) {
	syncMap := sync.Map{}

	for n := 0; n < b.N; n++ {
		syncMap.Store(n, n)
	}

}

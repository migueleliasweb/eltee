package benchmarks

import (
	"sync"
	"testing"
	"time"
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

//BenchmarkIntSyncMapAppend Benchmarks append on maps
func BenchmarkIntSyncMapAppend(b *testing.B) {
	syncMap := sync.Map{}

	now := time.Now().Unix()

	for n := 0; n < b.N; n++ {
		syncMap.Store(now, n)
	}

}

//BenchmarkIntSyncMapAppendParallel Benchmarks append on maps
func BenchmarkIntSyncMapAppendParallel(b *testing.B) {
	syncMap := sync.Map{}
	resultChan := make(chan bool)

	gorountinesCount := 8

	//match real "hashing" complexity
	now := time.Now().Unix()

	for g := 0; g < gorountinesCount; g++ {
		go func() {
			for n := 0; n < (b.N / gorountinesCount); n++ {
				syncMap.Store(now, n)
			}

			resultChan <- true
		}()
	}

	b.StartTimer()
	<-resultChan
	<-resultChan
	<-resultChan
	<-resultChan
	<-resultChan
	<-resultChan
	<-resultChan
	<-resultChan
	b.StopTimer()

}

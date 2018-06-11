package benchmarks

import (
	"testing"
)

//BenchmarkIntChannelAppend
func BenchmarkIntChannelAppend(b *testing.B) {
	intsChannel := make(chan int, b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		intsChannel <- 100
	}
	b.StopTimer()
}

//BenchmarkIntChannelDrain
func BenchmarkIntChannelDrain(b *testing.B) {
	intsChannel := make(chan int, b.N)

	for n := 0; n < b.N; n++ {
		intsChannel <- 100
	}

	close(intsChannel)

	b.StartTimer()
	for range intsChannel {
		//no op
	}
	b.StopTimer()
}

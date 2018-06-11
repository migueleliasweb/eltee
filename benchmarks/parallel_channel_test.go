package benchmarks

import "testing"

func channelAppend(inChan chan int, outChan chan bool, total int) {
	for n := 0; n < total; n++ {
		inChan <- 100
	}

	outChan <- true
}

//BenchmarkIntChannelParallelAppend
func BenchmarkIntChannelParallelAppend(b *testing.B) {
	intsChannel := make(chan int, b.N)
	resultChans := make(chan bool)

	for n := 0; n < 8; n++ {
		go func() {
			channelAppend(
				intsChannel,
				resultChans,
				b.N/8,
			)
		}()
	}

	b.StartTimer()
	<-resultChans
	<-resultChans
	<-resultChans
	<-resultChans
	<-resultChans
	<-resultChans
	<-resultChans
	<-resultChans
	b.StopTimer()
}

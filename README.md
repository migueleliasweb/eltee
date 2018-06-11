# ELTEE
(LT) ELTEE stands for "LOADTEST"

# Benchmarks

##### Running all the beanchmaks at once
```bash
$ go test -bench=. -v ./...
```
##### Some results
```bash
BenchmarkIntChannelAppend-8           	50000000	        26.6 ns/op
BenchmarkIntChannelDrain-8            	30000000	        53.7 ns/op
BenchmarkIntSliceAppend-8             	300000000	         3.91 ns/op
BenchmarkIntMapAppend-8               	10000000	       151 ns/op
BenchmarkIntSyncMapAppend-8           	20000000	       106 ns/op
BenchmarkIntSyncMapAppendParallel-8   	10000000	       198 ns/op
BenchmarkGoNumStdDev5000-8            	  100000	     12121 ns/op
BenchmarkGoNumMean5000-8              	  200000	      6127 ns/op
BenchmarkMathMin5000-8                	  100000	     22812 ns/op
BenchmarkMathMax5000-8                	  100000	     23532 ns/op
BenchmarkIntChannelParallelAppend-8   	30000000	        49.2 ns/op
```

##### Insights from the benchmarks

* Slices are fast
* Map are sinigicantly slower than slices (around 7 times slower)
* With increasing number os keys, maps tend not to change acess time too much for same amount of keys
* Key length and type matters to maps because of the hashing function used internally
* Map from sync.Map is a somewhat good option when parallel **write** access is required
* Channels are surprisingly fast for both append and drains
* Math functions can we quite slow for specifically Min & Max are substancially slower than first imagined
*

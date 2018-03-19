# Sort Benchmarking Results

>> go test -bench=.
goos: darwin
goarch: amd64
BenchmarkStableSort10-8             	2000000000	         0.00 ns/op
BenchmarkStableSort100-8            	2000000000	         0.00 ns/op
BenchmarkStableSort1000-8           	2000000000	         0.00 ns/op
BenchmarkStableSort10000-8          	2000000000	         0.00 ns/op
BenchmarkStableSort100000-8         	1000000000	         0.06 ns/op
BenchmarkStableSort1000000-8        	2000000000	         0.33 ns/op
BenchmarkStableSort10000000-8       	       1	7430872078 ns/op
BenchmarkNotInPlaceSort1-8          	2000000000	         0.00 ns/op
BenchmarkNotInPlaceSort10-8         	2000000000	         0.00 ns/op
BenchmarkNotInPlaceSort100-8        	2000000000	         0.00 ns/op
BenchmarkNotInPlaceSort1000-8       	2000000000	         0.00 ns/op
BenchmarkNotInPlaceSort10000-8      	2000000000	         0.00 ns/op
BenchmarkNotInPlaceSort100000-8     	2000000000	         0.01 ns/op
BenchmarkNotInPlaceSort1000000-8    	2000000000	         0.08 ns/op
BenchmarkNotInPlaceSort10000000-8   	       1	1561622314 ns/op
BenchmarkMergeSort1-8               	2000000000	         0.00 ns/op
BenchmarkMergeSort10-8              	2000000000	         0.00 ns/op
BenchmarkMergeSort100-8             	2000000000	         0.00 ns/op
BenchmarkMergeSort1000-8            	2000000000	         0.00 ns/op
BenchmarkMergeSort10000-8           	2000000000	         0.00 ns/op
BenchmarkMergeSort100000-8          	2000000000	         0.01 ns/op
BenchmarkMergeSort1000000-8         	2000000000	         0.09 ns/op
BenchmarkMergeSort10000000-8        	       1	2028252351 ns/op
PASS
ok  	_/Users/nerdy.wink/projects/allocations/go/src/sort	34.509s
## reset
    
Recommendation is to use the method in benchmark1
    
    
    >  ./reset 
    a @ p=0x1400000e030 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    b @ p=0x1400000e018 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    [0 0 0 0 0 0 0 0 0 0 10 0 0 0 0 0 0 0 0 0 20 0 0 0 0 0 0 0 0 0 0 0]
    a @ p=0x1400000e030 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    
    
### benchmark1
    
    benchmark1 >  go test -bench=. -count 5
    goos: darwin
    goarch: arm64
    pkg: github.com/rpcox/gobox/slice/reset/benchmark1
    cpu: Apple M4 Pro
    BenchmarkReset-14    	618404888	         1.779 ns/op
    BenchmarkReset-14    	682849730	         1.766 ns/op
    BenchmarkReset-14    	685212286	         1.766 ns/op
    BenchmarkReset-14    	675054058	         1.778 ns/op
    BenchmarkReset-14    	677525265	         1.774 ns/op
    PASS
    ok  	github.com/rpcox/gobox/slice/reset/benchmark1	6.151s
    
    
### benchmark2
    
    benchmark2 >  go test -bench=. -count 5
    goos: darwin
    goarch: arm64
    pkg: github.com/rpcox/gobox/slice/reset/benchmark2
    cpu: Apple M4 Pro
    BenchmarkReset-14    	 4634409	       241.1 ns/op
    BenchmarkReset-14    	 5068075	       236.5 ns/op
    BenchmarkReset-14    	 5075872	       236.3 ns/op
    BenchmarkReset-14    	 5077761	       236.2 ns/op
    BenchmarkReset-14    	 5065082	       236.3 ns/op
    PASS
    ok  	github.com/rpcox/gobox/slice/reset/benchmark2	6.140s
    

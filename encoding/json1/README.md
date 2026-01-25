## json test

Go std lib vs jasoniter

    >  go test -bench=.
    goos: darwin
    goarch: arm64
    pkg: github.com/rpcox/bench-test/cmd/json1
    cpu: Apple M4 Pro
    BenchmarkMarshalStd-14      	 6722593	       157.6 ns/op
    BenchmarkMarshalJsi-14      	 5576372	       214.9 ns/op
    BenchmarkUnMarshalStd-14    	 1988365	       604.7 ns/op
    BenchmarkUnMarshalJsi-14    	 8841778	       133.8 ns/op
    PASS
    ok  	github.com/rpcox/bench-test/cmd/json1	5.052s



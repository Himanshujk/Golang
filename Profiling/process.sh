//for cpu testing
go test -cpuprofile cpu.prof -bench .
go tool pprof cpu.prof

//memory testing
go test -memprofile mem.prof -bench .
go tool pprof mem.prof

//for traces
go test -cpuprofile cpu.prof -bench .
go test -trace trace.out
go tool trace trace.out

//flamegraph
go test -cpuprofile cpu.prof -bench .
go tool pprof -http=":8081" test.test.exe cpu.prof
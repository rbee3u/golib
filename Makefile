test:
	go test -v -cover -coverprofile=cover.out ./...

test_bloomfilter:
	go test -v -cover -coverprofile=cover.out ./bloomfilter/...

test_ibch:
	go test -v -cover -coverprofile=cover.out ./ibch/...

test_memo:
	go test -v -cover -coverprofile=cover.out ./memo/...

test_runner:
	go test -v -cover -coverprofile=cover.out ./runner/...

coverout:
	go tool cover -html=cover.out

bench_bloomfilter:
	go test -v -bench=. -run=^$$ -benchtime=10s -cpuprofile=cpu.out -benchmem -memprofile=mem.out ./bloomfilter

bench_ibch:
	go test -v -bench=. -run=^$$ -benchtime=10s -cpuprofile=cpu.out -benchmem -memprofile=mem.out ./ibch

bench_memo:
	go test -v -bench=. -run=^$$ -benchtime=10s -cpuprofile=cpu.out -benchmem -memprofile=mem.out ./memo

cpuout:
	go tool pprof -http=: cpu.out

memout:
	go tool pprof -http=: mem.out

lint:
	golangci-config-generator
	golangci-lint run

install-gcg:
	go install github.com/rbee3u/golangci-config-generator/cmd/golangci-config-generator@latest

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

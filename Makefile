.PHONY: bench
bench:
	go test --run XXX --bench -cpuprofile cpu.prof -memprofile mem.prof ./core
.PHONY: pprof
pprof: 
	go tool pprof -http 0.0.0.0:8080 ./cpu.prof

.PHONY: proto
proto:
	protoc -I. \
	--go_out=. \
	--go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	--proto_path=. \
	proto/*.proto


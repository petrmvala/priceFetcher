build:
	go build -o bin/pricefetcher

run: build
	./bin/pricefetcher

image:
	docker build -t fetcher:1 --platform linux/amd64 .

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: proto

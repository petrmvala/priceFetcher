build:
	go build -o bin/pricefetcher

run: build
	./bin/pricefetcher

image:
	docker build -t fetcher:1 --platform linux/amd64 .

# Price Fetcher Microservice

API fetching service for cryptocurrency prices.

* Logging and metrics decoration pattern, separated in files
* API Error handling in one place
* API client included

## Development and running

Run or build locally via Makefile (`make run`), or build a Docker image (`make image`). Then run the image:

```
docker run -p 3000:3000 --platform=linux/amd64 -d fetcher:1
```

### Protobuf

On MacOS, install protobuf via Homebrew:

```
brew install protobuf
```

GRPC and Protobuffer package dependencies:

```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u google.golang.org/grpc
go install google.golang.org/grpc
```

Add `protoc-gen-go-grpc` to your PATH:

```
export PATH=$PATH:$(go env GOPATH)/bin
```

## Testing

Manually for example with curl or browser.

Successful test

```
❯ curl 'http://localhost:3000/?ticker=BTC'
{"ticker":"BTC","price":20000}
```

Unsuccessful test

```
❯ curl 'http://localhost:3000/?ticker=BT'
{"error":"the given ticker (BT) is not supported"}
```

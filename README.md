### Price Fetcher Microservice

API fetching service for cryptocurrency prices.

* Logging and metrics decoration pattern, separated in files
* API Error handling in one place
* API client included

### Development and running

Run or build locally via Makefile (`make run`), or build a Docker image (`make image`). Then run the image:

```
docker run -p 3000:3000 --platform=linux/amd64 -d fetcher:1
```

### Testing

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

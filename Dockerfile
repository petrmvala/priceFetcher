FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod .

COPY . ./

RUN go build -o /pricefetcher

EXPOSE 3000

CMD [ "/pricefetcher" ]

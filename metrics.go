package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	// push to Prometheus or so
	fmt.Println("Pushing metrics to Prometheus")
	return s.next.FetchPrice(ctx, ticker)
}

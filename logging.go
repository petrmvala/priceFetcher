package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type LoggingService struct {
	next   PriceFetcher
	logger slog.Logger
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	return &LoggingService{
		next:   next,
		logger: *logger,
	}
}

func (s *LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		s.logger.Info("fetchPrice",
			slog.Any("requestID", ctx.Value("requestID")),
			slog.Float64("price", price),
			slog.Duration("took_ns", time.Since(begin)),
			slog.Any("err", err),
		)
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

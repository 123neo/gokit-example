package endpoint

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
const (
	LoggerContext = "logger"
)

func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Info("endpoint_logging_middleware", zap.NamedError("transport_error", err),
					zap.Duration("took", time.Since(begin)))

			}(time.Now())
			ctx = context.WithValue(ctx, LoggerContext, logger)
			return next(ctx, request)

		}
	}
}

package demo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"ultra-hand/pkg/log"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Info("msg", "calling endpoint")
			defer func(logger log.Logger, keyVals ...interface{}) {
				logger.Info("msg", keyVals)
			}(logger, "msg", "called endpoint")
			return next(ctx, request)
		}
	}
}

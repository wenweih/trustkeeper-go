package endpoint

import (
	"context"
	"fmt"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"

	// httptransport "github.com/go-kit/kit/transport/http"
	stdjwt "github.com/go-kit/kit/auth/jwt"
)

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// AuthMiddleware returns an endpoint middleware
// 添加 middleware , 减少每次 decodeXXXRequest 都要从 request header 中 extract Authorization 参数
func AuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// Add your middleware logic here
			if _, ok := ctx.Value(stdjwt.JWTTokenContextKey).(string); !ok {
				// 若需要鉴权的接口请求头中不携带 Authorization jwt token, 直接给客户端返回错误，不请求鉴权服务
				return nil, fmt.Errorf("Authorization Token Empty")
			}
			return next(ctx, request)
		}
	}
}

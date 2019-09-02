package endpoint

import (
	"context"
	"fmt"
	"time"
	"github.com/afex/hystrix-go/hystrix"
	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"
	// "github.com/go-kit/kit/circuitbreaker"
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

// Hystrix returns an endpoint.Middleware that implements the circuit
// breaker pattern using the afex/hystrix-go package.
//
// When using this circuit breaker, please configure your commands separately.
//
// See https://godoc.org/github.com/afex/hystrix-go/hystrix for more
// information.
// https://www.ru-rocker.com/2017/04/24/micro-services-using-go-kit-hystrix-circuit-breaker/
// https://www.twblogs.net/a/5b85da1e2b71775d1cd439b0
func Hystrix(commandName string, fallbackMesg string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
				Timeout: 1000,
				MaxConcurrentRequests: 1000,
			})
			var resp interface{}
			if err := hystrix.Do(commandName, func() (err error) {
				resp, err = next(ctx, request)
				return err
			}, func(err error) error {
				logger.Log("fallbackErrorDesc", err.Error())
				resp = struct {
					Fallback string `json:"fallback"`
				}{
					fallbackMesg,
				}
				return nil
			}); err != nil {
				return nil, err
			}
			return resp, nil
		}
	}
}

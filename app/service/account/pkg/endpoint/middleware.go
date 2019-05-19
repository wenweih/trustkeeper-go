package endpoint

import (
	"context"
	"fmt"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"
	service "trustkeeper-go/app/service/account/pkg/service"
	"github.com/dgrijalva/jwt-go"
	"trustkeeper-go/app/service/account/pkg/configure"
)

type Claims struct {
	jwt.StandardClaims
}

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
func AuthMiddleware(conf configure.Conf, s service.AccountService) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			claims := &Claims{}
			req := request.(SignoutRequest)
			tkn, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(conf.JWTKey), nil
			})

			if err != nil || !tkn.Valid {
				return nil, fmt.Errorf(err.Error())
			}

			acc ,err := s.FindByTokenID(ctx, claims.Id)
			if err != nil {
				return nil, fmt.Errorf("token was reset" + err.Error())
			}
			return next(context.WithValue(ctx, "account", acc), request)
		}
	}
}

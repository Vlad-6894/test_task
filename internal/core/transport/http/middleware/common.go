package core_http_middleware

import (
	"net/http"
	"time"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_response "github.com/Vlad-6894/test_task/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDKey  = "X-Request-ID"
	originKey     = "Origin"
	AccesOrigin   = "Access-Control-Allow-Origin"
	AccessMethods = "Access-Control-Allow-Methods"
	AccessHeaders = "Access-Control-Allow-Headers"
)

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := map[string]struct{}{
				"http://localhost:5050": {},
			}

			origin := r.Header.Get(originKey)

			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set(AccesOrigin, origin)
				w.Header().Set(AccessMethods, "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set(AccessHeaders, "Content-Type, Authorization")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDKey)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDKey, requestID)
			w.Header().Set(requestIDKey, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(logger *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDKey)

			log := logger.With(zap.String("request_id", requestID), zap.String("URL", r.URL.String()))

			ctx := core_logger.ToContext(r.Context(), log)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "Panic in the programm!")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			beforeTime := time.Now()
			log.Info(
				"incoming HTTP request",
				zap.String("http_method", r.Method),
				zap.Time("time", beforeTime.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Info(
				"done HTTP request",
				zap.Int("status code: ", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(beforeTime)),
			)
		})
	}
}

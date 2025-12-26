package middlewares

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Example metric: count of requests
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// Example metric: request duration
	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func PrometheusMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			status := c.Response().Status
			method := c.Request().Method
			endpoint := c.Path()

			// Increment request counter
			requestsTotal.WithLabelValues(method, endpoint, fmt.Sprintf("%d", status)).Inc()

			// Record request duration
			duration := time.Since(start).Seconds()
			requestDuration.WithLabelValues(method, endpoint).Observe(duration)

			return err
		}
	}
}

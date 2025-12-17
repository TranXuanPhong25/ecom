package routes

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricRoute(e *echo.Echo) {
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}

package telexgin

import (
	"runtime"
	"time"

	"github.com/Micah-Shallom/go-apm-sdk/telex"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RequestMetricsCollector(startTime time.Time, c *gin.Context) telex.Metrics {
	latency := time.Since(startTime)
	statusCode := c.Writer.Status()

	//request count metrics

	requestMetrics := telex.Metrics{
		RequestMetrics: telex.RequestMetrics{
			Path:       c.Request.URL.Path,
			Method:     c.Request.Method,
			Latency:    latency.String(),
			StatusCode: statusCode,
		},
	}

	return requestMetrics

}

func (h *Handler) PerformanceMetricsCollector() telex.Metrics{
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	performanceMetrics := telex.Metrics{
		PerformanceMetrics: telex.PerformanceMetrics{
			MemoryAllocation: memStats.Alloc,
			CPUUsage:         runtime.NumCPU(),
			MemoryUsage:      memStats.Sys,
			NetworkTraffic:   0,
			GCCycles:         memStats.NumGC,
			Goroutines:       runtime.NumGoroutine(),
		},
	}
	return performanceMetrics
}

func (h *Handler) ErrorMetricsCollector() telex.Metrics{

	//error count metrics

	errorMetrics := telex.Metrics{
		ErrorMetrics: telex.ErrorMetrics{},
	}
	return errorMetrics

}

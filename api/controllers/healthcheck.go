package controllers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/linkly/db"
)

const (
	StatusOK                 string = "OK"
	StatusPartiallyAvailable string = "Partially Available"
	StatusUnavailable        string = "Unavailable"
	StatusTimeout            string = "Timeout during health check"
)

type (
	// Check represents the health check response.
	Check struct {
		// Status is the check status.
		Status string `json:"status"`
		// Timestamp is the time in which the check occurred.
		Timestamp string `json:"timestamp"`
		// StartUp is the time to boot up the system.
		StartUp string `json:"startup"`
		// Uptime is the time in which the check occurred.
		Uptime string `json:"uptime"`
		// Failures holds the failed checks along with their messages.
		Failures map[string]string `json:"failures,omitempty"`
		// System holds information of the go process.
		System `json:"system"`
	}
	// System runtime variables about the go process.
	System struct {
		// Version is the go version.
		Version string `json:"version"`
		// GoroutinesCount is the number of the current goroutines.
		GoroutinesCount int `json:"goroutines_count"`
		// TotalAllocBytes is the total bytes allocated.
		TotalAllocBytes int `json:"total_alloc_bytes"`
		// HeapObjectsCount is the number of objects in the go heap.
		HeapObjectsCount int `json:"heap_objects_count"`
		// TotalAllocBytes is the bytes allocated and not yet freed.
		AllocBytes int `json:"alloc_bytes"`
	}
)

// NewSystemMetrics returns a new SystemMetrics struct
func NewSystemMetrics() System {
	s := runtime.MemStats{}
	runtime.ReadMemStats(&s)

	return System{
		Version:          runtime.Version(),
		GoroutinesCount:  runtime.NumGoroutine(),
		TotalAllocBytes:  int(s.TotalAlloc),
		HeapObjectsCount: int(s.HeapObjects),
		AllocBytes:       int(s.Alloc),
	}
}

// NewCheck returns a new HealthCheck struct
func NewCheck(status string, failures map[string]string, startTime time.Time, bootTime time.Duration) Check {
	return Check{
		Status:    status,
		Timestamp: time.Now().Format(time.RFC3339),
		StartUp:   bootTime.String(),
		Uptime:    time.Since(startTime).String(),
		Failures:  failures,
		System:    NewSystemMetrics(),
	}
}

type HealthCheck interface {
	HealthCheck(ctx *gin.Context, startTime time.Time, bootTime time.Duration)
}

type healthCheck struct {
}

// HealthCheck godoc
// @Summary Show the status of the system.
// @Description checks the health of the system.
// @ID healthcheck
// @Tags health
// @Produce json
// @Success 200 {object} controllers.Check
// @Router /api/health [get]
func (m *healthCheck) HealthCheck(ctx *gin.Context, startTime time.Time, bootTime time.Duration) {
	status := StatusUnavailable
	failures := make(map[string]string)

	if db.IsConnected {
		status = StatusOK
	} else {
		status = StatusPartiallyAvailable
		failures["postgres"] = "failed during postgresql health check"
	}

	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.JSON(http.StatusOK, NewCheck(
		status,
		failures,
		startTime,
		bootTime,
	))
}

func NewHealthCheck() HealthCheck {
	return &healthCheck{}
}

package collector

import (
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/time/rate"
)

const (
	namespace string = "crtsh"
	subsystem string = "exporter"
)

var (
	// Rate Limiter is shared across collectors
	ratelimiter = rate.NewLimiter(rate.Every(time.Second), 4)
)

func BuildFQName(name string, logger *slog.Logger) string {
	logger.Info("Creating Metric",
		"name", name,
	)
	return prometheus.BuildFQName(namespace, subsystem, name)
}

package collector

import (
	"fmt"
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

type Domain struct {
	name string
	l    int
}

func NewDomain(name string) (*Domain, error) {
	if name == "" {
		return nil, fmt.Errorf("domain name cannot be empty")
	}
	return &Domain{
		name: name,
		l:    len(name),
	}, nil
}
func (d *Domain) Hostname(fqName string) (string, error) {
	l := len(fqName)
	if l == 0 {
		return "", fmt.Errorf("fqName is empty")
	}
	if l < d.l+1 {
		return "", fmt.Errorf("fqName is too short; expected '{host}.%s'", d.name)
	}
	return fqName[:len(fqName)-(d.l+1)], nil
}

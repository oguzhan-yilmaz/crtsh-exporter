package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/DazWilkin/crtsh-exporter/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	sRoot string = `
<h2>A Prometheus Exporter for <a href="https://crt.sh">Certificate Search</a></h2>
<ul>
	<li><a href="{{ .Metrics }}">metrics</a></li>
	<li><a href="/healthz">healthz</a></li>
</ul>`
)

var (
	// GitCommit is the git commit value and is expected to be set during build
	GitCommit string
	// GoVersion is the Golang runtime version
	GoVersion = runtime.Version()
	// OSVersion is the OS version (uname --kernel-release) and is expected to be set during build
	OSVersion string
	// StartTime is the start time of the exporter represented as a UNIX epoch
	StartTime = time.Now().Unix()
)
var (
	endpoint    = flag.String("endpoint", ":8080", "The endpoint of the Exporter HTTP server")
	hostList    = flag.String("hosts", "", "Comma-separated list of hosts")
	metricsPath = flag.String("path", "/metrics", "The path on which Prometheus metrics will be served")
)
var (
	tRoot = template.Must(template.New("root").Parse(sRoot))
)

func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w)
	tRoot.Execute(w, struct {
		Metrics string
	}{
		Metrics: *metricsPath,
	})
}
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if GitCommit == "" {
		logger.Info("GitCommit value unchanged: expected to be set during build")
	}
	if OSVersion == "" {
		logger.Info("OSVersion value unchanged: expected to be set during build")
	}

	flag.Parse()

	if *hostList == "" {
		logger.Info("Expect '--hosts' to contain at least one host")
	}

	hosts := strings.Split(*hostList, ",")

	registry := prometheus.NewRegistry()
	registry.MustRegister(collector.NewExporterCollector(OSVersion, GoVersion, GitCommit, StartTime))
	registry.MustRegister(collector.NewHostsCollector(hosts, logger))

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))
	mux.Handle("/healthz", http.HandlerFunc(handleHealthz))
	mux.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	logger.Info("Server starting",
		"endpoint", *endpoint,
	)
	logger.Info("metrics path",
		"path", *metricsPath,
	)
	logger.Error("Server failed", http.ListenAndServe(*endpoint, mux))
}

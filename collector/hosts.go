package collector

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	baseURL string = "https://crt.sh"
)

type Record struct {
	IssueCAID      int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int    `json:"id"`
	EntryTimestamp Time   `json:"entry_timestamp"`
	NotBefore      Time   `json:"not_before"`
	NotAfter       Time   `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}

type Time struct {
	time.Time
}

const expiryDateLayout = "2006-01-02T15:04:05"

func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(expiryDateLayout, s)
	return
}

// HostsCollector collects metrics
type HostsCollector struct {
	client *http.Client
	hosts  []string
	log    *slog.Logger

	// Metrics
	Certificate *prometheus.Desc
}

// NewHostsCollector is a function that returns a new HostCollector
func NewHostsCollector(hosts []string, log *slog.Logger) *HostsCollector {
	log = log.With("function", "HostsCollector")
	client := &http.Client{}
	return &HostsCollector{
		client: client,
		hosts:  hosts,
		log:    log,

		Certificate: prometheus.NewDesc(
			BuildFQName("certificate", log),
			"Certificate count",
			[]string{
				"name",
				"not_before",
				"not_after",
				"serial_number",
			},
			nil,
		),
	}
}

// Collect implements Prometheus' Collector interface and is used to collect metrics
func (c *HostsCollector) Collect(ch chan<- prometheus.Metric) {
	log := c.log.With("method", "Collect")

	for _, host := range c.hosts {
		log := log.With("host", host)
		rqst, err := http.NewRequest(http.MethodGet, baseURL, nil)
		if err != nil {
			msg := "unable to create request"
			log.Info(msg)
			continue
		}

		q := rqst.URL.Query()
		q.Add("q", host)
		q.Add("output", "json")
		rqst.URL.RawQuery = q.Encode()

		resp, err := c.client.Do(rqst)
		if err != nil {
			msg := "unable to execute request"
			log.Info(msg)
			continue
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			msg := "unable to read response body"
			log.Info(msg)
			continue
		}

		// Debugging
		// log.Info("Response",
		// 	"body", string(b),
		// )

		records := []Record{}
		if err := json.Unmarshal(b, &records); err != nil {
			msg := "unable to unmarshal response"
			log.Info(msg)
			continue
		}

		log.Info("Records",
			"number", len(records),
		)
		if len(records) == 0 {
			msg := "expected at least one record in response"
			log.Info(msg)
			continue
		}

		// Grab most recent entry
		record := records[0]

		log.Info("Most recent record",
			"record", record,
		)
		ch <- prometheus.MustNewConstMetric(
			c.Certificate,
			prometheus.GaugeValue,
			float64(len(records)),
			[]string{
				record.NameValue,
				strconv.FormatInt(record.NotBefore.Unix(), 10),
				strconv.FormatInt(record.NotAfter.Unix(), 10),
				record.SerialNumber,
			}...,
		)
	}
}

// Describe implements Prometheus' Collector interface and is used to describe metrics
func (c *HostsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Certificate
}

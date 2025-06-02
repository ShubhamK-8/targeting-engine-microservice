package prometheous

import (
	"github.com/prometheus/client_golang/prometheus"
)

// --- Prometheus Metrics ---

var (
	// Request counters
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"path", "method", "status"},
	)

	// Request duration histogram
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets, // Default buckets for common latencies
		},
		[]string{"path", "method"},
	)

	// Redis cache hit/miss counter
	RedisCacheHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_cache_hits_total",
			Help: "Total number of Redis cache hits.",
		},
	)
	RedisCacheMisses = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_cache_misses_total",
			Help: "Total number of Redis cache misses.",
		},
	)

	// Elasticsearch query duration histogram
	EsQueryDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "elasticsearch_query_duration_seconds",
			Help:    "Duration of Elasticsearch queries.",
			Buckets: prometheus.DefBuckets,
		},
	)

	// Campaigns returned gauge (could also be summary or histogram if needed more granularity)
	CampaignsReturned = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "campaigns_returned_count",
			Help: "Number of campaigns returned in the last successful response.",
		},
	)
)

// --- Initialization of Prometheus Metrics ---

func InitPrometheus() {
	prometheus.MustRegister(HttpRequestsTotal)
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(RedisCacheHits)
	prometheus.MustRegister(RedisCacheMisses)
	prometheus.MustRegister(EsQueryDuration)
	prometheus.MustRegister(CampaignsReturned)
}

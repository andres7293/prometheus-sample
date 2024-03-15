package main

import (
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_memory_usage",
		Help: "Memory usage of my app",
	})
	httpRequests = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_http_requests",
		Help: "Number of http requests",
	})
)

func init() {
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(httpRequests)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Set memory usage gauge value
	memoryUsage.Set(float64(m.Alloc))
	httpRequests.Inc()
	w.Write([]byte("Hello, world!"))
}

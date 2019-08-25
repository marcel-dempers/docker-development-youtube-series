package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsProcessed = promauto.NewCounter(prometheus.CounterOpts{
			Name: "request_operations_total",
			Help: "The total number of processed requests",
	})
)

var (
    requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name:    "request_duration_seconds",
        Help:    "Histogram for the runtime of a simple example function.",
        Buckets: prometheus.LinearBuckets(0.5, 1.0, 2),
    })
)



func main() {
	
	fmt.Println("starting...")
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		
		timer := prometheus.NewTimer(requestDuration)
		defer timer.ObserveDuration()
		time.Sleep(1000 * time.Millisecond) //Sleep for a second
		fmt.Fprint(w, "Welcome to my application!")
		
		requestsProcessed.Inc()
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":80", nil)	

}

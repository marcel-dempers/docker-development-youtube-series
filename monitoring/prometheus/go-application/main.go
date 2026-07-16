package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var	requestsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_request_operations_total",
		Help: "The total number of processed requests",
	})

var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "go_request_duration_seconds",
	Help:    "Histogram for the duration in seconds.",
	Buckets: []float64{1, 2, 5, 6, 10},
	}	,
	[]string{"endpoint"},
)

func main() {

	fmt.Println("starting...")

	prometheus.MustRegister(requestDuration)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		//start a timer
		start := time.Now()

		time.Sleep(600 * time.Millisecond) //Sleep simulate work
		fmt.Fprint(w, "Welcome to my application!")
	
		//measure the duration and log to prometheus
		httpDuration := time.Since(start)
		requestDuration.WithLabelValues("GET /").Observe(httpDuration.Seconds())
		
		//increment a counter for number of requests processed
		requestsProcessed.Inc()
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)

}
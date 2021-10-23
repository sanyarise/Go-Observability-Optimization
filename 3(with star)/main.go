package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func server1()  error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe("localhost:8001", mux)
}

func server2() error {
	mux2 := http.NewServeMux()
	mux2.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe("localhost:8002", mux2)
}

func main() {
	
	usersRegistered := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "users_registered",
		})
	prometheus.MustRegister(usersRegistered)

	usersOnline := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "users_online",
		})
	prometheus.MustRegister(usersOnline)

	requestProcessingTimeSummaryMs := prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name:       "request_processing_time_summary_ms",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		})
	prometheus.MustRegister(requestProcessingTimeSummaryMs)

	requestProcessingTimeHistogramMs := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "request_processing_time_histogram_ms",
			Buckets: prometheus.LinearBuckets(0, 10, 20),
		})
	prometheus.MustRegister(requestProcessingTimeHistogramMs)

	go func() {
		for {
			usersRegistered.Inc()
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			for i := 0; i < 10000; i++ {
				usersOnline.Set(float64(i))
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	go func() {
		src := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(src)
		for {
			obs := float64(100 + rnd.Intn(30))
			requestProcessingTimeSummaryMs.Observe(obs)
			requestProcessingTimeHistogramMs.Observe(obs)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		log.Printf("Server1 starting")
		err1 := server1()
		if err1 != nil {
			log.Fatal("Fatal error on port 8001")
		}
	}()
	

	go func() {
		log.Printf("Server2 starting")
		err := server2()
		if err != nil {
		log.Fatal("Fatal error on port 8002")
		}
	}()

	select{}
}

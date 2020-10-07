package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var (
	hashesGenerated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "butte_hashes_generated_total",
		Help: "The total number of hashes generated",
	})
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":2112", nil)
	}()

	for true {
		_, err := bcrypt.GenerateFromPassword([]byte{'f', 'i', 'r', 'e'}, 10)
		if err != nil {
			panic(err)
		}
		hashesGenerated.Inc()
	}
}

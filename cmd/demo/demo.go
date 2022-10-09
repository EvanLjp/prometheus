package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	codeStatus := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "status_code_total",
		Help: "help status_code total desc",
	}, []string{"code"})

	go func() {
		for {
			codeStatus.WithLabelValues("200").Add(1)
			time.Sleep(time.Second)
		}
	}()
	registry := prometheus.NewRegistry()
	registry.MustRegister(codeStatus)
	// 暴露
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8000", nil)
}

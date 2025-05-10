package main

import (
	"github.com/konnenl/load-balancer/internal/balancer"
	"github.com/konnenl/load-balancer/internal/config"
	"github.com/konnenl/load-balancer/internal/logger"
	"net/http"
)

func main() {
	logger := logger.New()

	loader := config.NewLoader("json")
	cfg, err := loader.Load("config.json")
	if err != nil {
		logger.ErrorLog.Fatal("Failed to load config:", err)
	}
	var servers []*balancer.Server
	for _, s := range cfg.Servers {
		servers = append(servers, &balancer.Server{
			Url: s.Url,
		})
	}

	balancer := balancer.New(cfg.Algorithm, servers, logger)

	http.HandleFunc("/", balancer.HandleRequest)
	port := ":" + cfg.Port
	logger.InfoLog.Printf("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		logger.ErrorLog.Fatal(err)
	}
}

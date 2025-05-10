package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/konnenl/load-balancer/internal/balancer"
	"github.com/konnenl/load-balancer/internal/config"
)

func main() {
	loader := config.NewLoader("json")
	cfg, err := loader.Load("config.json")
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
	var servers []*balancer.Server
	for _, s := range cfg.Servers {
		servers = append(servers, &balancer.Server{
			Url: s.Url,
		})
	}

	balancer := balancer.New(cfg.Algorithm, servers)

	http.HandleFunc("/", balancer.HandleRequest)
	port := ":" + cfg.Port
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/konnenl/load-balancer/internal/config"
	"github.com/konnenl/load-balancer/internal/balancer"
	"net/http"
	"fmt"
	"log"
)

func main() {  
	loader := config.NewLoader("json")
	cfg, err := loader.Load("config.json")
	if err != nil{
		log.Fatal("Config error: %v", err)
	}
	var servers []*balancer.Server
	for _, s := range cfg.Servers{
		servers = append(servers, &balancer.Server{
			Url: s.Url,
		})
	}

	balancer := balancer.New("round-robin", servers)

	http.HandleFunc("/", balancer.HandleRequest)
	port := ":" + cfg.Port
	fmt.Println("Started on :", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/konnenl/load-balancer/internal/config"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v\n", cfg)
	})


	port := ":" + cfg.Port
	fmt.Println("Started on :", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

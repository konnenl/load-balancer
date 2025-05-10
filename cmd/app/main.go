package main

import (
	"github.com/konnenl/load-balancer/internal/balancer"
	"github.com/konnenl/load-balancer/internal/config"
	"github.com/konnenl/load-balancer/internal/logger"
	"net/http"
)

func main() {
	// Инициализация логгера
	logger := logger.New()

	// Инициализация загрузчика конфига в json формате
	loader := config.NewLoader("json")
	// Загрузка конфига
	cfg, err := loader.Load("config.json")
	if err != nil {
		logger.ErrorLog.Fatal("Failed to load config:", err)
	}

	// Преобразование информации о серверах в массив структур Server
	var servers []*balancer.Server
	for _, s := range cfg.Servers {
		servers = append(servers, &balancer.Server{
			Url: s.Url,
		})
	}

	// Инициализация балансировщика с алгоритмом, указанным в конфиге
	balancer := balancer.New(cfg.Algorithm, servers, logger)

	// Регистрация обработчика для всех входящих HTTP запросов
	http.HandleFunc("/", balancer.HandleRequest)
	port := ":" + cfg.Port
	// Запуск HTTP сервера на порте, указанном в конфиге
	logger.InfoLog.Printf("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		logger.ErrorLog.Fatal(err)
	}
}

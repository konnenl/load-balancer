package balancer

import (
	"github.com/konnenl/load-balancer/internal/logger"
	"net/http"
	"sync"
)

// Структура сервера, на который перенаправляются запросы
type Server struct {
	Url string
	mux sync.RWMutex
}

// Интерфейс для структур алгоритмов балансировки
type Balancer interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

// Функция, возвращающая реализацию балансировщика в зависимости от выбранного алгоритма
func New(algotithm string, servers []*Server, logger *logger.Logger) Balancer {
	switch algotithm {
	case "round-robin":
		return NewRoundRobinBalancer(servers, logger)
	default:
		return NewRoundRobinBalancer(servers, logger)
	}

}

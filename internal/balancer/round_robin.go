package balancer

import (
	"github.com/konnenl/load-balancer/internal/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type RoundRobinBalancer struct {
	servers []*Server
	current uint
	proxy   *httputil.ReverseProxy
	mux     sync.Mutex
	logger  *logger.Logger
}

func NewRoundRobinBalancer(servers []*Server, logger *logger.Logger) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		servers: servers,
		proxy:   &httputil.ReverseProxy{},
		logger:  logger,
	}
}

func (b *RoundRobinBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	b.logger.RequestLog.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
	nextServer := b.GetNext()
	if nextServer == nil {
		b.logger.InfoLog.Println("No available servers")
		http.Error(w, "Server unavailable", http.StatusServiceUnavailable)
		return
	}
	url, _ := url.Parse(nextServer.Url)
	b.proxy.Director = func(r *http.Request) {
		r.URL.Scheme = url.Scheme
		r.URL.Host = url.Host
	}

	b.logger.InfoLog.Printf("Request redirected to %s", nextServer.Url)
	b.proxy.ServeHTTP(w, r)
}

func (b *RoundRobinBalancer) GetNext() *Server {
	b.mux.Lock()
	defer b.mux.Unlock()
	l := len(b.servers)
	for i := 0; i < l; i++ {
		b.current = (b.current + 1) % uint(l)

		server := b.servers[b.current]

		server.mux.RLock()
		alive := b.IsAlive(server.Url)
		server.mux.RUnlock()

		if alive {
			return server
		}
		b.logger.InfoLog.Printf("Server  unavailable: %s", server.Url)
	}
	return nil
}

func (b *RoundRobinBalancer) IsAlive(url string) bool {
	client := http.Client{Timeout: 1 * time.Second}
	_, err := client.Head(url)
	return err == nil
}

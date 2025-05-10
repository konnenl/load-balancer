package balancer

import(
	"net/url"
	"fmt"
	"sync"
	"net/http"
	"net/http/httputil"
	"time"
)

type RoundRobinBalancer struct{
	servers []*Server
	current uint
	proxy *httputil.ReverseProxy
	mux sync.Mutex
}

func NewRoundRobinBalancer(servers []*Server) *RoundRobinBalancer{
	return &RoundRobinBalancer{
		servers: servers,
		proxy: &httputil.ReverseProxy{},
	}
}

func(b *RoundRobinBalancer) HandleRequest(w http.ResponseWriter, r *http.Request){
	nextServer := b.GetNext()
	if nextServer == nil{
		http.Error(w, "No available servers", 509)
		return 
	}
	url, _ := url.Parse(nextServer.Url)
	b.proxy.Director = func(r *http.Request){
		r.URL.Scheme = url.Scheme
		r.URL.Host = url.Host
	}
	b.proxy.ServeHTTP(w, r)
}

func (b *RoundRobinBalancer) GetNext() *Server{
	b.mux.Lock()
	defer b.mux.Unlock()
	l := len(b.servers)
	for i := 0; i < l; i++{
		b.current = (b.current + 1) % uint(l)

		server := b.servers[b.current]

		server.mux.RLock()
		alive := b.IsAlive(server.Url)
		server.mux.RUnlock()
		
		if alive{
			fmt.Println("Selected", alive, server.Url)
			return server
		}
		fmt.Println("Skipped", alive, server.Url)
	}

	return nil
}

func (b *RoundRobinBalancer) IsAlive(url string) bool{
	client := http.Client{Timeout: 1 * time.Second}
    _, err := client.Head(url)
    return err == nil
}
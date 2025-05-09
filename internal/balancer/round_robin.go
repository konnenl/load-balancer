package balancer

import(
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type RoundRobinBalancer struct{
	servers []*Server
	current uint
	proxy *httputil.ReverseProxy
	lock sync.Mutex
}

func NewRoundRobinBalancer(servers []*Server) *RoundRobinBalancer{
	return &RoundRobinBalancer{
		servers: servers,
		proxy: &httputil.ReverseProxy{},
	}
}

func(b *RoundRobinBalancer) HandleRequest(w http.ResponseWriter, r *http.Request){
	nextServer := b.GetNext()
	url, _ := url.Parse(nextServer.Url)
	b.proxy.Director = func(r *http.Request){
		r.URL.Scheme = url.Scheme
		r.URL.Host = url.Host
	}
	b.proxy.ServeHTTP(w, r)
}

func (b *RoundRobinBalancer) GetNext() *Server{
	b.lock.Lock()
	defer b.lock.Unlock()
	b.current = (b.current + 1) % uint(len(b.servers))
	return b.servers[b.current]
}
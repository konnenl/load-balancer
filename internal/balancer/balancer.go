package balancer

import(
	"net/http"
)

type Server struct{
	Url string
}

type Balancer interface{
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

func New(algotithm string, servers []*Server) Balancer{
	switch algotithm {
	case "round-robin":
		return NewRoundRobinBalancer(servers)
	default:
		return NewRoundRobinBalancer(servers)
	}

}
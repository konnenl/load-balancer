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
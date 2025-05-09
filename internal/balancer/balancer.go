package balancer

func New(algotithm string, servers []*Server) Balancer{
	switch algotithm {
	case "round-robin":
		return NewRoundRobinBalancer(servers)
	default:
		return NewRoundRobinBalancer(servers)
	}

}
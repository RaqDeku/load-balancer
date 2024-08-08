package internals

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	Port            string
	roundRobinCount int
	servers         []Server
}

func CreateLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvialableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsHealthy() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvialableServer()

	fmt.Printf("forwarding request to %q\n", targetServer.Address())

	targetServer.Serve(rw, r)
}

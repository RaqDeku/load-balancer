package internals

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	Address() string

	IsHealthy() bool

	Serve(rw http.ResponseWriter, r *http.Request)
}

type NodeServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

func (s *NodeServer) Address() string {
	return s.address
}

func (s *NodeServer) IsHealthy() bool {
	return true
}

func (s *NodeServer) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw, r)
}

func CreateNodeServer(addr string) *NodeServer {
	serverUrl, err := url.Parse(addr)

	if err != nil {
		fmt.Print(err)
	}

	return &NodeServer{
		address: addr,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/load-balancer/internals"
)

func main() {
	servers := []internals.Server{
		internals.CreateNodeServer("https://bing.com"),
		internals.CreateNodeServer("https://google.com"),
	}

	lb := internals.CreateLoadBalancer("8080", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.ServeProxy(rw, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}

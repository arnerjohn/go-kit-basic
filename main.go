package main

import (
	"github.com/arnerjohn/go-kit-basic/service"
	"github.com/arnerjohn/go-kit-basic/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

var message = "Hello World"

func main() {
	var svc service.ServiceInterface
	svc = service.Service{}

	uppercaseHandler := httptransport.NewServer(
		transport.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		transport.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

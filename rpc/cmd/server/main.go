package main

import (
	"github.com/pPrecel/go-examples/rpc/pkg/math"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/rpc"
)

const address = ":1234"

func main() {
	// ### RPC CONFIGURATION
	val := new(math.Server)
	err := rpc.Register(val)
	failOnErr(err)

	rpc.HandleHTTP()

	// ### LISTEN AND SERVE
	log.Printf("Listen and serve %s...\n", address)
	http.ListenAndServe(address, nil)

}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

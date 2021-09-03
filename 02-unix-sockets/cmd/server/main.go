package main

import (
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/pPrecel/go-examples/unix-sockets/pkg/math"
	log "github.com/sirupsen/logrus"
)

func main() {
	// ### CONFIGURATION
	log.Infof("Remove old socket: %s", math.Address)
	err := os.RemoveAll(math.Address)
	failOnErr(err)

	// ### RPC CONFIGURATION
	val := new(math.Server)
	err = rpc.Register(val)
	failOnErr(err)

	rpc.HandleHTTP()

	// ### LISTEN AND SERVE
	log.Info("Listen...")
	listener, err := net.Listen(math.Network, math.Address)
	failOnErr(err)
	defer listener.Close()

	log.Infof("Serve new %s socket: %s\n", math.Network, math.Address)
	http.Serve(listener, nil)
}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

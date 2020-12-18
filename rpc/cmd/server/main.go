package main

import (
	"github.com/pPrecel/rpc-test/pkg/math"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/rpc"
)


func main() {
	val := new(math.Server)
	err := rpc.Register(val)
	failOnErr(err)

	rpc.HandleHTTP()

	log.Printf("Listen: localhost:1234\n")
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	log.Println("Serve...")
	http.Serve(l, nil)

}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

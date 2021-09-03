package main

import (
	"time"

	"github.com/pPrecel/go-examples/rpc/pkg/math"
	log "github.com/sirupsen/logrus"
)

func main() {
	a := 5
	b := 20
	start := time.Now()
	log.Printf("5 + 20 = %d\n", a+b)
	log.Printf("Duration: %s", time.Until(start).String())

	start = time.Now()
	if val, err := math.Plus(a, b); err == nil {
		log.Printf("5 + 20 = %d\n", val)
		log.Printf("Duration: %s", time.Until(start).String())
	} else {
		log.Fatal(err)
	}
}

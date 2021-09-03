package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/pPrecel/go-examples/port-scanner/pkg/parse"
	log "github.com/sirupsen/logrus"
)

const (
	portsFlagName       = "ports"
	addressFlagName     = "address"
	workersFlagName     = "workers"
	defaultPorts        = "1-1024"
	defaultAddress      = "scanme.nmap.org"
	defaultWorkers      = 100
	targetAddressFormat = "%s:%d"
)

func main() {
	log.Println("Load configuration...")
	address := flag.String(addressFlagName, defaultAddress, "Target address")
	ports := flag.String(portsFlagName, defaultPorts, "Target ports")
	workers := flag.Int(workersFlagName, defaultWorkers, "Amount of workers")
	flag.Parse()

	portsChan := make(chan int, *workers)
	resultChan := make(chan int)

	log.Printf("Running %d workers...", *workers)
	for i := 0; i < *workers; i++ {
		go runWorker(*address, portsChan, resultChan)
	}

	log.Printf("Scanning ports: %s, on address: %s...", *ports, *address)
	targetPorts, err := parse.Range(*ports)
	if err != nil {
		log.Panic(err)
	}

	go func(targetPorts []int) {
		for _, port := range targetPorts {
			portsChan <- port
		}
	}(targetPorts)

	for _ = range targetPorts {
		port := <-resultChan
		if port == 0 {
			continue
		}

		log.Printf("Port %d is open", port)
	}

	close(portsChan)
	close(resultChan)
	log.Println("Done :)")
}

func runWorker(address string, ports, result chan int) {
	for port := range ports {
		conn, err := net.Dial("tcp", fmt.Sprintf(targetAddressFormat, address, port))
		if err != nil {
			result <- 0
		} else {
			conn.Close()
			result <- port
		}
	}
}

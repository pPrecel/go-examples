package main

import (
	"fmt"
	"os"

	"github.com/pPrecel/go-examples/shodan/pkg/shodan"
)

const (
	shodanKeyEnv = "SHODAN_API_KEY"
	infoFormat   = "Plan: %s\nQueryCredits: %d\nScanCredits: %d\n\n"
	hostFormat   = "IPString: %s:%d \nLocation: %s %s\n\n"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: 'go run cmd/main.go [one_argument]'")
	}

	APIKey := os.Getenv(shodanKeyEnv)
	if APIKey == "" {
		fmt.Printf("Missing token env... (%s)\n", shodanKeyEnv)
		os.Exit(1)
	} else {
		fmt.Printf("Starting with token: %s...\n", APIKey[0:4])
	}

	c := shodan.New(APIKey)
	info, err := c.APIInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Client info:")

	fmt.Printf(infoFormat, info.Plan, info.QueryCredits, info.ScanCredits)

	hosts, err := c.Hosts(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Hosts:")
	for _, host := range hosts.Matches {
		fmt.Printf(hostFormat, host.IPString, host.Port,
			host.Location.CountryCode, host.Location.City)
	}
}

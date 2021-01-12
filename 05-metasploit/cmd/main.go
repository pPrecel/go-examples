package main

import (
	"os"

	"github.com/pPrecel/go-examples/metasploit/pkg/msf"
	log "github.com/sirupsen/logrus"
)

const (
	hostEnvName     = "METASPLOIT_SERVER_HOST"
	usernameEnvName = "METASPLOIT_SERVER_USERNAME"
	passwordEnvName = "METASPLOIT_SERVER_PASSWORD"
)

func main() {
	log.Info("Read configuration...")
	host := os.Getenv(hostEnvName)
	user := os.Getenv(usernameEnvName)
	pass := os.Getenv(passwordEnvName)

	if host == "" || user == "" || pass == "" {
		log.Fatalf("Please set envs: %s, %s, %s", hostEnvName, usernameEnvName, passwordEnvName)
	}

	log.Info("Connecting...")
	client, err := msf.New(host, user, pass)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Logout()

	log.Info("List session...")
	res, err := client.SessionList()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Server version: %+v", res)
}

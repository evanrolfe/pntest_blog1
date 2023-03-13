package main

import (
	"fmt"
	"log"

	"github.com/docker/go-plugins-helpers/network"
	mydriver "github.com/tugbadartici/docker-network-plugin"
)

//------------------------------------------------------------------------------
// build:
//------------------------------------------------------------------------------
// docker build -t my-plugin-image .
// docker run --name my-network-plugin -v /var/run/docker.sock:/var/run/docker.sock my-plugin-image

// docker network create --driver my-network-plugin --subnet 172.18.0.0/16 --gateway 172.18.0.1 my-network

func main() {
	fmt.Println("Starting...")
	d := mydriver.NewDriver()
	h := network.NewHandler(d)

	err := h.ServeUnix("my-network", 0)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}

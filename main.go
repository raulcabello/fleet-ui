package main

import (
	"github.com/raulcabello/fleet-ui/internal/k8s"
	"github.com/raulcabello/fleet-ui/internal/server"
	"log"
)

func main() {
	client, err := k8s.NewClient()
	if err != nil {
		log.Fatalln(err.Error())
	}

	httpServer := server.NewHttp(client)
	httpServer.Start()

}

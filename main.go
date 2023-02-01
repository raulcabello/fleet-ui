package main

import (
	"github.com/raulcabello/fleet-ui/internal/client"
	"github.com/raulcabello/fleet-ui/internal/server"
	"log"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		log.Fatalln(err.Error())
	}

	httpServer := server.NewHttp(client)
	httpServer.Start()

}

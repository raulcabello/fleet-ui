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
	/*	fmt.Println(client.GetBundleList("fleet-default"))
		fmt.Println(client.GetBundle("fleet-default", "test-simple"))
		fmt.Println(client.GetGitRepoList("fleet-default"))
	*/
	//	fmt.Println(fc.Fleet().V1alpha1().Bundle().Cache().List("", labels.Everything()))
}

package main

import (
	"github.com/Napigo/go-account-service/frameworks/rest"
	"github.com/Napigo/go-account-service/interfaces"

	"github.com/Napigo/npgc"
)

func main() {
	npgc.Load(".env")

	// List of all servers need to initialzied and run
	// such as Rest API, grpc, Websocket etc..
	servers := []interfaces.Framework{
		rest.RestServer{},
	}

	for _, server := range servers {
		server.Run()
	}
	var restServer interfaces.Framework = rest.RestServer{}
	restServer.Run()
}

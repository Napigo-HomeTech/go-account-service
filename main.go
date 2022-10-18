package main

import (
	commonrest "github.com/Napigo/go-account-service/frameworks/rest"
	"github.com/Napigo/go-account-service/interfaces"
	"github.com/Napigo/npgcommon"
)

func main() {
	npgcommon.LoadEnv()

	var restServer interfaces.Framework = commonrest.RestServer{}
	restServer.Run()
}

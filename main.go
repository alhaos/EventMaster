package main

import "github.com/alhaos/EventMaster/apiServer"

func main() {

	r := apiServer.New()

	r.Run(":8080")

}

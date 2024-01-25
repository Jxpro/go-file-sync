package main

import (
	"flag"
	"fmt"
	"github.com/Jxpro/go-file-sync/client"
	"github.com/Jxpro/go-file-sync/common"
	"github.com/Jxpro/go-file-sync/server"
)

func main() {
	args := common.ParseArgs()
	switch args["mode"] {
	case "client":
		fmt.Println("Starting client...")
		client.Start()
	case "server":
		fmt.Println("Starting server...")
		server.Start()
	default:
		fmt.Println("Invalid mode: " + args["mode"])
		flag.Usage()
	}
}

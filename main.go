package main

import (
	"flag"
	"fmt"
	"github.com/Jxpro/go-file-sync/client"
	"github.com/Jxpro/go-file-sync/common"
	"github.com/Jxpro/go-file-sync/server"
)

func main() {
	args := common.GetArgs()
	switch args.Mode {
	case common.ClientMode:
		fmt.Println("Starting client...")
		client.Start()
	case common.ServerMode:
		fmt.Println("Starting server...")
		server.Start()
	default:
		fmt.Println("Invalid mode: " + args.Mode)
		flag.Usage()
	}
}

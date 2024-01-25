package main

import (
	"flag"
	"fmt"
	"github.com/Jxpro/go-file-sync/client"
	"github.com/Jxpro/go-file-sync/server"
)

var (
	dir    string
	key    string
	mode   string
	addr   string
	config string
)

func init() {
	flag.StringVar(&dir, "d", "./data", "directory to sync")
	flag.StringVar(&dir, "dir", "./data", "directory to sync")

	flag.StringVar(&key, "k", "goFileSync12138", "authentication key")
	flag.StringVar(&key, "key", "goFileSync12138", "authentication key")

	flag.StringVar(&mode, "m", "server", "start mode: server or client")
	flag.StringVar(&mode, "mode", "server", "start mode: server or client")

	flag.StringVar(&addr, "a", "127.0.0.1:6880", "connect or listen address")
	flag.StringVar(&addr, "addr", "127.0.0.1:6880", "connect or listen address")

	// 配置文件的优先级低于命令行参数
	flag.StringVar(&config, "c", "./config.json", "config file path")
	flag.StringVar(&config, "config", "./config.json", "config file path")

	flag.Usage = func() {
		fmt.Println("Usage: go-file-sync [options]")
		fmt.Println("Options:")
		fmt.Println("  -d, --dir      directory to sync, default: ./data")
		fmt.Println("  -k, --key      authentication key, default: goFileSync12138")
		fmt.Println("  -m, --mode     start mode: server or client, default: server")
		fmt.Println("  -a, --addr     address to connect or listen, default:127.0.0.1:6880")
		fmt.Println("  -c, --config   config file path (least priority), default: ./config.json")
		fmt.Println("  -h, --help     show this help message")
	}
}

func main() {
	// 解析命令行参数
	flag.Parse()
	switch mode {
	case "client":
		fmt.Println("Starting client...")
		client.Start()
	case "server":
		fmt.Println("Starting server...")
		server.Start()
	default:
		fmt.Println("Invalid mode: " + mode)
		flag.Usage()
	}
}

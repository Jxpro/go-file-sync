package common

import (
	"flag"
	"fmt"
)

func ParseArgs() map[string]string {
	// 声明命令行参数
	var dir, key, mode, addr, config string

	// 绑定命令行参数
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

	// 自定义帮助信息
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

	// 解析命令行参数
	flag.Parse()

	// 包装命令行参数
	args := map[string]string{
		"dir":    dir,
		"key":    key,
		"mode":   mode,
		"addr":   addr,
		"config": config,
	}

	// 返回命令行参数
	return args
}

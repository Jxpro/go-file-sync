package common

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sync"
)

type Arguments struct {
	// 使用标签自动解析json格式的配置文件
	Dir    string `json:"dir"`
	Key    string `json:"key"`
	Mode   string `json:"mode"`
	Addr   string `json:"addr"`
	Config string `json:"args"`
}

const (
	ClientMode = "client"
	ServerMode = "server"
)

var (
	once     sync.Once
	instance *Arguments
)

func GetArgs() *Arguments {
	// 单例模式，只解析一次命令行参数和配置文件
	// 不使用init函数是因为在测试中存在错误
	once.Do(func() {
		instance = &Arguments{
			Dir:  "./data",
			Key:  "goFileSync12138",
			Mode: "server",
			Addr: "127.0.0.1:6880",
		}
		// 解析命令行参数和配置文件
		commandArgs := parseCommand()
		fileArgs := parseFile(commandArgs.Config)

		// 合并参数，命令行参数优先级高于配置文件
		instance.mergeArgs(&fileArgs)
		instance.mergeArgs(&commandArgs)

		// 打印参数
		instance.printArgs()
	})
	return instance
}

func parseCommand() (commandArgs Arguments) {
	// 绑定命令行参数
	flag.StringVar(&commandArgs.Dir, "d", commandArgs.Dir, "directory to sync")
	flag.StringVar(&commandArgs.Dir, "dir", commandArgs.Dir, "directory to sync")

	flag.StringVar(&commandArgs.Key, "k", commandArgs.Key, "authentication key")
	flag.StringVar(&commandArgs.Key, "key", commandArgs.Key, "authentication key")

	flag.StringVar(&commandArgs.Mode, "m", commandArgs.Mode, "start mode: server or client")
	flag.StringVar(&commandArgs.Mode, "mode", commandArgs.Mode, "start mode: server or client")

	flag.StringVar(&commandArgs.Addr, "a", commandArgs.Addr, "connect or listen address")
	flag.StringVar(&commandArgs.Addr, "addr", commandArgs.Addr, "connect or listen address")

	// 配置文件的优先级低于命令行参数
	flag.StringVar(&commandArgs.Config, "c", "", "args file path")
	flag.StringVar(&commandArgs.Config, "args", "", "args file path")

	// 自定义帮助信息
	flag.Usage = func() {
		fmt.Println("Usage: go-file-sync [options]")
		fmt.Println("Options:")
		fmt.Println("  -d, --dir      directory to sync, default: ./data")
		fmt.Println("  -k, --key      authentication key, default: goFileSync12138")
		fmt.Println("  -m, --mode     start mode: server or client, default: server")
		fmt.Println("  -a, --addr     address to connect or listen, default:127.0.0.1:6880")
		fmt.Println("  -c, --args   args file path, priority less than command line parameters")
		fmt.Println("  -h, --help     show this help message")
	}

	// 解析命令行参数
	flag.Parse()
	return
}

func parseFile(path string) (fileArgs Arguments) {
	// 如果配置文件路径为空，则不解析
	if path == "" {
		return
	}

	// 读取配置文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("------ Error ------\n" + err.Error())
		os.Exit(1)
	}

	// 延迟关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("------ Error ------\n" + err.Error())
			os.Exit(1)
		}
	}(file)

	// 解析配置文件
	if err := json.NewDecoder(file).Decode(&fileArgs); err != nil {
		fmt.Printf("Error decoding args file: %s\n", err)
	}
	return
}

func (a *Arguments) mergeArgs(new *Arguments) {
	// 使用反射遍历结构体，合并参数
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		// 获取字段名
		key := v.Type().Field(i).Name
		// 获取新参数的字段值
		value := reflect.ValueOf(new).Elem().FieldByName(key)
		// 如果字段值不为空，则更新
		if value.String() != "" {
			v.Field(i).Set(value)
		}
	}
}

func (a *Arguments) printArgs() {
	fmt.Println("Initializing with arguments:")
	// 使用反射遍历结构体，打印参数
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		// 获取字段名
		key := v.Type().Field(i).Name
		// 获取字段值
		value := v.Field(i).String()
		fmt.Printf("  %s: %s\n", key, value)
	}
}

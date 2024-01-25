package common

import (
	"os"
	"testing"
)

func TestArgs(t *testing.T) {
	// 测试命令行参数
	os.Args = []string{"go-file-sync", "-c", "config_test.json", "-d", "./data"}
	args := GetArgs()
	// 测试单例模式
	args = GetArgs()

	switch {
	// 命令行参数优先级最高
	case args.Dir != "./data":
		t.Error("Wrong dir")
	// 命令行和配置文件均为包含的参数使用默认值
	case args.Key != "goFileSync12138":
		t.Error("Wrong key")
	// 无命令行参数且有配置文件时使用配置文件提供的参数
	case args.Mode != "server":
		t.Error("Wrong mode")
	case args.Addr != "0.0.0.0:8080":
		t.Error("Wrong addr")
	}
}

func TestMergeArgs(t *testing.T) {
	// 测试合并参数
	args := &Arguments{
		Dir:  "./test",
		Key:  "test",
		Mode: "client",
		Addr: "0.0.0.0:8080",
	}
	args.mergeArgs(&Arguments{
		Dir:  "./dataNew",
		Key:  "goFileSync12138",
		Mode: "server",
		Addr: "127.0.0.1:6880",
	})

	switch {
	case args.Dir != "./dataNew":
		t.Error("Wrong dir")
	case args.Key != "goFileSync12138":
		t.Error("Wrong key")
	case args.Mode != "server":
		t.Error("Wrong mode")
	case args.Addr != "127.0.0.1:6880":
		t.Error("Wrong addr")
	}
}

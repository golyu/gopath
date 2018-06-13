package main

import (
	"os"
	"path/filepath"
	"flag"
	"strings"
	"fmt"
	"os/exec"
	"log"
)

var (
	// 需要设置为临时go path的路径
	curPath string
)

func main() {
	//flag.StringVar(&curPath, "path", "", "要设置为动态path的路径,不设置则为当前路径,如果是在src下,就设置为src的上级目录")
	//if curPath == "" {
	// 当前路径
	pwd, _ := os.Getwd()
	curPath = pwd
	if strings.HasSuffix(pwd, "/src") {
		curPath = strings.TrimSuffix(pwd, "/src")
	}
	// 上级目录,也就是如果你在 /User/aaa/bbb/ccc/src/ddd下执行,我也会把/User/aaa/bbb/ccc设置为环境变量
	upPath := strings.TrimSuffix(pwd, "/"+filepath.Base(pwd))
	if strings.HasSuffix(upPath, "/src") {
		curPath = strings.TrimSuffix(upPath, "/src")
	}

	//}
	// 设置GoPath
	ensureGoPath()

	// 对应命令执行
	cmds := flag.Args()
	if len(cmds) > 2 {
		runPrint(cmds[1], cmds[1:])
	}
}

// ensureGoPath 配置 GoPath
func ensureGoPath() {
	// 全局设置的GoPath
	rootPath := os.Getenv("GOPATH")
	// 当前路径
	var fullPath string
	if rootPath == "" {
		fullPath = curPath
	} else {
		fullPath = rootPath + ":" + curPath
	}
	err := os.Setenv("GOPATH", fullPath)
	if err != nil {
		fmt.Println("环境变量设置失败")
	} else {
		fmt.Println("当前命令行的临时环境变量为:", fullPath)
	}
}

// runPrint 执行命令
func runPrint(cmd string, args []string) {
	eCmd := exec.Command(cmd, args...)
	eCmd.Stdout = os.Stdout
	eCmd.Stderr = os.Stderr
	err := eCmd.Run()
	if err != nil {
		log.Fatal("执行命令发生错误:", err.Error())
	}

}

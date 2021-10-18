package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/net/webdav"
)

// 获取参数
func getAgs() (fullAddr string, path string) {

	var dir = flag.String("p", ".", "共享路径")
	var addr = flag.String("a", "", "地址")
	var port = flag.Int("port", 8080, "端口")

	flag.Parse()

	var argCount = len(flag.Args())

	if argCount == 0 && *dir == "." {
		flag.Usage()
		os.Exit(0)
	}

	if argCount == 1 {
		arg1 := flag.Args()[0]

		if reflect.TypeOf(arg1).String() == "string" {
			*dir = arg1
		} else {
			flag.Usage()
			os.Exit(0)
		}
	}

	if *addr == "" {
		fullAddr = ":" + strconv.Itoa(*port)
	} else {
		fullAddr = *addr + ":" + strconv.Itoa(*port)
	}

	path = *dir
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func main() {

	addr, path := getAgs()
	p, err := PathExists(path)

	if !p {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	fmt.Println("WebDav Sever run ...")
	fmt.Printf("Run as http://%s\n", addr)
	fmt.Printf("Run directory %s\n", path)

	err = http.ListenAndServe(addr, &webdav.Handler{
		FileSystem: webdav.Dir(path),
		LockSystem: webdav.NewMemLS(),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

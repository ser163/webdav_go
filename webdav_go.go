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
func getAgs() (fullAddr string,
	path string,
	sslFlag bool,
	keyFile string,
	certFile string,
	userName string,
	passWord string,
	readMode bool,
) {

	var dir = flag.String("p", ".", "共享路径")
	var addr = flag.String("a", "", "地址")
	var port = flag.Int("port", 8080, "端口")
	var ssl = flag.Bool("ssl", false, "https: true http: false")
	var key = flag.String("ssl-key", "key.pem", "https key file")
	var cert = flag.String("ssl-cert", "cert.pem", "https cert file")
	var user = flag.String("user", "", "user name")
	var pass = flag.String("pass", "", "password")
	var readOnly = flag.Bool("read", false, "read only (defalut: false)")

	flag.Parse()

	var argCount = len(os.Args)
	if argCount == 1 && *dir == "." {
		flag.Usage()
		os.Exit(0)
	}

	if argCount == 2 {
		arg1 := os.Args[1]
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
	sslFlag = *ssl
	keyFile = *key
	certFile = *cert
	userName = *user
	passWord = *pass
	readMode = *readOnly

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

	addr, path, sslMode, keyFile, certFile, user, pass, readMode := getAgs()
	// fmt.Println(addr, path, sslMode, keyFile, certFile, user, pass, readMode)
	// 判断目录是否存在
	p, err := PathExists(path)

	if !p {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	fmt.Println("WebDav Sever run ...")
	var sslStr string
	if sslMode {
		sslStr = "https://"
	} else {
		sslStr = "http://"
	}
	fmt.Printf("Run as %s%s\n", sslStr, addr)
	fmt.Printf("Run directory %s\n", path)

	fs := &webdav.Handler{
		FileSystem: webdav.Dir(path),
		LockSystem: webdav.NewMemLS(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(user, pass)
		if user != "" && pass != "" {
			fmt.Println("Authenticate Check")
			username, password, ok := req.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if username != user || password != pass {
				http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
				return
			}
		}
		if readMode {
			switch req.Method {
			case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
				http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
				return
			}
		}
		// fmt.Println(req)
		fs.ServeHTTP(w, req)
	})

	// 判断是否是ssl模式
	if sslMode {
		err = http.ListenAndServeTLS(addr, certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(addr, nil)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	//获取当前目录
	var host string

	//解析命令行参数
	flag.StringVar(&host, "host", "127.0.0.1:1234", "监听地址")
	flag.Parse()

	//获取当前路径
	dir, _ := os.Getwd()
	//设置静态文件目录
	router.ServeFiles("/*filepath", http.Dir(dir))
	fmt.Println(fmt.Sprintf("当前所在目录：%s", dir))

	log.Printf("server start at http://%s", host)
	err := http.ListenAndServe(host, router)
	if err != nil {
		log.Fatal(err)
	}
}

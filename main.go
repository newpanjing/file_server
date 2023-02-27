package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	//获取当前目录
	pwd, _ := os.Getwd()

	var host string

	//解析命令行参数
	flag.StringVar(&host, "host", "127.0.0.1:1234", "监听地址")
	flag.Parse()

	//设置静态文件目录
	router.ServeFiles("/*filepath", http.Dir(pwd))
	log.Printf("server start at http://%s", host)
	_ = http.ListenAndServe(host, router)
}

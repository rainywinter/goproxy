package main

import (
	"flag"
	"github.com/rainywinter/goproxy"
	"log"
	"net/http"
	"fmt"
)

var (
	port int
	socksAddr string
)

func init(){
	flag.IntVar(&port,"p",8080,"http proxy server listen port")
	flag.StringVar(&socksAddr,"socks","127.0.0.1:1080","socks5 server addr")
}

func main(){
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer(socksAddr)
	addr := fmt.Sprintf(":%d",port)
	log.Printf("server start,listening %s\n",addr)
	log.Fatal(http.ListenAndServe(addr,proxy))
}

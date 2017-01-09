# Introduction

Modify from [https://github.com/elazarl/goproxy](https://github.com/elazarl/goproxy)

A Http Proxy that forward http/https protocol to socks5.  
Expecially for command line usage,like curl and wget &. 

# Usage

		go get github.com/rainywinter/goproxy/goproxy
		
		goproxy -p 8080 -socks 127.0.0.1:1080

>create a http proxy server,listen 8080 port,send data through localhost socks5 server.

Example:

		http_proxy=localhost:8080 https_proxy=localhost:8080 wget https://github.com
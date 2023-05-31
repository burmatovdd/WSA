package main

import "WAF_Analytics/internal/server"

func main() {
	service := server.Service{}
	service.HandleRequest()
}

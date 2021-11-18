package main

import (
	"github.com/emersonluiz/go-email/server"
	"github.com/emersonluiz/go-email/service"
)

func main() {
	service.ReceiveMessage()
	server := server.NewServer()
	server.Run()
}

package main

import "github.com/emersonluiz/go-user/server"

func main() {
	server := server.NewServer()
	server.Run()
}

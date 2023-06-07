package main

import (
	"github.com/aabdullahgungor/go-restapi-redis/server"
)

func main() {

	s := server.NewServer()
	s.Run()
}

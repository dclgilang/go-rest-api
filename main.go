package main

import (
	"github.com/dclgilang/go-rest-api/database"
	"github.com/dclgilang/go-rest-api/server"
)

func main() {
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}

package main

import (
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatal("Unable to connect to mongo")
	}
	defer ms.Close()

	u := mongo.NewGfUserService(ms.Copy(), "buyright", "user")
	s := server.NewServer(u)

	s.Start()
}
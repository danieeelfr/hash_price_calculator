package main

import (
	"log"

	srv "github.com/danieeelfr/hash_price_calculator/internal/controller/grpc_server"
	rep "github.com/danieeelfr/hash_price_calculator/internal/repository/dbprovider"
)

func main() {

	db, e := rep.NewDBProvider()
	if e != nil {
		log.Panic(e)
	}
	// u := db.GetUser("1")
	print(db)

	s, e := srv.NewGRPCServer()
	if e != nil {
		log.Panic(e)
	}

	s.Start()

}

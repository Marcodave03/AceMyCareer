package main

import (
	"github.com/KRook0110/PSD/rentaltool/backend/api"
	"log"
)

func main() {
	db, err := api.NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	db.Init()
	server := api.NewApiServer(":8080", db)

	server.Run()
}

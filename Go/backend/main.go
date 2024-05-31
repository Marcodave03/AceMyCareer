package main

import (
	"fmt"

	"github.com/KRook0110/PSD/rentaltool/backend/api"
)

func main() {

    apiServer := api.CreateApiServer(":8080")
    fmt.Println("listening on", "localhost:8080")

    apiServer.Run()
}

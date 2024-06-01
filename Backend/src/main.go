package main

import(

    "github.com/Marcodave03/AceMyCareer/backend/src/api"

)

func main() {
    apiserver := api.CreateNewApiServer("localhost:8080")
    apiserver.Run()
}

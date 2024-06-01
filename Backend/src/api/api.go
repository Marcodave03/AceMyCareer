package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/lib/pq"
)

type ApiServer struct {
	ListenAddr string
	db             *sql.DB
}

func CreateDB() *sql.DB {
	connStr := "user=postgres dbname=postgres sslmode=disable password=PasswordYangSusah port=8092"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateNewApiServer(listenPort string) *ApiServer {
	return &ApiServer{
		ListenAddr: listenPort,
		db:             CreateDB(),
	}
}

func (s *ApiServer) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		WriteJson(w, "kontol", http.StatusOK)
	})

    fmt.Println("Api Server Listening on ", s.ListenAddr)
    log.Fatal(http.ListenAndServe(s.ListenAddr, router))

}

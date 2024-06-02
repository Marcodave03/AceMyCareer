package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Marcodave03/AceMyCareer/backend/src/api/users"
	"github.com/Marcodave03/AceMyCareer/backend/src/api/utils"
)

type ApiServer struct {
	ListenAddr string
	db         *sql.DB
}

func CreateDB() *sql.DB {
	connStr := "user=postgres dbname=postgres sslmode=disable password=PasswordYangSusah port=8092" // TODO: should use a config file
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateNewApiServer(listenPort string) *ApiServer {
	return &ApiServer{
		ListenAddr: listenPort,
		db:         CreateDB(),
	}
}

func (s *ApiServer) setupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { utils.WriteJson(w, "titit", http.StatusOK) })

	// Users
	userHandler := users.CreateUserHandler(s.db)
	router.HandleFunc("/api/users", userHandler.HandleUsers)
	router.HandleFunc("/api/users/{username}", userHandler.HandleUserByUsername)

	// Static File Servers
	// router.Handle("/api/images/profilepicture/", http.StripPrefix("/api/images/profilepicture/", http.FileServer(http.Dir(""))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(os.Getenv("API_STATIC_FILES_DIRECTORY")))))

	return router
}

func (s *ApiServer) Run() {
	// Loading env files
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Api Server Listening on ", s.ListenAddr)
	log.Fatal(http.ListenAndServe(s.ListenAddr, s.setupRoutes()))
}

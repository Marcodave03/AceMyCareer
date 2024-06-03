package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"database/sql"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Marcodave03/AceMyCareer/backend/src/api/interviews"
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

func (s *ApiServer) handleUploadImages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:

		err := r.ParseMultipartForm(10 << 24) // 10 MB Max
		if err != nil {
            http.Error(w, "Form Too Big", http.StatusBadRequest)
            return
		}

        file, header, err := r.FormFile("image")
		if err != nil {
            http.Error(w, "Unable to read form", http.StatusBadRequest)
            return
		}
        defer file.Close()

        filename := strings.Replace(uuid.New().String() + filepath.Ext(header.Filename), "-", "", -1)
        filepath := filepath.Join(os.Getenv("API_STATIC_FILES_DIRECTORY"), filename)

        dst, err := os.Create(filepath)
        if err != nil {
            http.Error(w, "Unable to create file", http.StatusInternalServerError)
            return
        }
        defer dst.Close()

        if _, err := io.Copy(dst, file); err != nil {
            http.Error(w, "Unable to create file", http.StatusInternalServerError)
            return
        }

        if err := utils.WriteJson(w, fmt.Sprintf("Upload Succesfull : %s", filename), http.StatusOK); err != nil {
            http.Error(w, "Unable to create file", http.StatusInternalServerError)
            return
        }
		break

	default:
		http.Error(w, "Method not suppported", http.StatusMethodNotAllowed)
        return
	}
}

func (s *ApiServer) createAllTables(w http.ResponseWriter, r *http.Request) {

    if err := users.CreateTableUsers(s.db); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := interviews.CreateTableInterviews(s.db); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (s *ApiServer) setupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { utils.WriteJson(w, "titit", http.StatusOK) })

    // Utility
    router.HandleFunc("/api/utils/create_all_tables", s.createAllTables)

	// Users
	userHandler := users.CreateUserHandler(s.db)
	router.HandleFunc("/api/users", userHandler.HandleUsers)
	router.HandleFunc("/api/users/{username}", userHandler.HandleUserByUsername)

    // Levels
    interviewHandler := interviews.CreateInterviewHandler(s.db)
    router.HandleFunc("/api/levels", interviewHandler.HandleLevel)

	// Static File Servers
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(os.Getenv("API_STATIC_FILES_DIRECTORY"))))) // images
    router.HandleFunc("/api/images/",s.handleUploadImages)


	return router
}

func (s *ApiServer) Run() {
	// Loading env files
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Api Server Listening on ", s.ListenAddr)
	log.Fatal(http.ListenAndServeTLS(s.ListenAddr, "server.crt", "server.key",s.setupRoutes()))
}

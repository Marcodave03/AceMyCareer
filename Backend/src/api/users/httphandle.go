package users

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type userHandler struct {
	DB *sql.DB
}

func CreateUserHandler(db *sql.DB) *userHandler {
	return &userHandler{
		DB: db,
	}
}

// URL : api/users
func (s *userHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleUserGet(w, r, s.DB)
		break
	case http.MethodPost:
		handleUserPost(w, r, s.DB)
		break
	case http.MethodPatch:
		handleUserPatch(w, r, s.DB)
		break
	case http.MethodDelete:
		handleUserDelete(w, r, s.DB)
		break
	}
}

func handleUserGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

func handleUserPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

func handleUserPatch(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

func handleUserDelete(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

// URL : api/users/<username>
func (s *userHandler) HandleUserByUsername(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
        handleUserByUsernameGet(w, r, s.DB)
		break
	}
}

func handleUserByUsernameGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    targetUsername := mux.Vars(r)["username"]
}

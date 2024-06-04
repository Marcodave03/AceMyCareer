package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Marcodave03/AceMyCareer/backend/src/api/utils"
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

// Users {{{
// URL : api/users
func (s *userHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
    utils.EnableCors(&w)
	switch r.Method {
	case http.MethodGet:
		handleUserGet(w, r, s.DB) // gets all users
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
    allUsers, err := getAllUsersFromTableUser(db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := utils.WriteJson(w, allUsers, http.StatusOK); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func handleUserPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var userInformation User
    if err := json.NewDecoder(r.Body).Decode(&userInformation); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    userInTable, err := checkUserUsernameInUserTable(db, userInformation.Username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if userInTable { // if is already in database
        http.Error(w, "Username already in table", http.StatusBadRequest)
        return
    }

    if err := insertUser(db, userInformation); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    utils.WriteJson(w, "Insertion Succes", http.StatusOK)
}

func handleUserPatch(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

func handleUserDelete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var userCredentials UserCredentialRequest
    if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    valid, err:= checkUserCredentials(db, userCredentials)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return

    }
    if !valid {
        http.Error(w, "Password Invalid", http.StatusUnauthorized)
        return
    }

    if err := deleteUser(db, userCredentials.Username); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.WriteJson(w, "Delete User", http.StatusOK)
    return
}

// URL : api/users/<username>
func (s *userHandler) HandleUserByUsername(w http.ResponseWriter, r *http.Request) {
    utils.EnableCors(&w)
	switch r.Method {
	case http.MethodGet:
        handleUserByUsernameGet(w, r, s.DB)
		break
	}
}

func handleUserByUsernameGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    targetUsername := mux.Vars(r)["username"]

    founduser, err := getUserbyUsername(db, targetUsername)
    if err == sql.ErrNoRows {
        http.Error(w, "Not Found",http.StatusBadRequest)
        return
    }

    if err := utils.WriteJson(w,founduser, http.StatusOK); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}// }}}




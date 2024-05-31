package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type ApiServer struct {
	listenPort string
	store Storage
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error


// Creates a new api server instance
func NewApiServer(listenAddr string, store Storage) *ApiServer {
	return &ApiServer{
		listenPort: listenAddr,
		store: store,
	}
}


func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/account/{username}", makeHandlerFunc(s.handleAccount))
	// router.HandleFunc("/account", makeHandlerFunc(s.handleAccount))

	fmt.Println("Server is running on : ", s.listenPort)
	return http.ListenAndServe(s.listenPort, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "UPDATE" {
		return s.handleUpdateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w,r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
func (s * ApiServer) handleGetSalt(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	newAccount := makeAccount(CreateAccountRequest{

    })
	// vars := mux.Vars(r)
	// fmt.Println(vars["username"])

	return WriteJSON(w, http.StatusOK, newAccount )
}
func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *ApiServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	userAccountInformation :=  new(CreateAccountRequest);
	err := json.NewDecoder(r.Body).Decode(userAccountInformation)
	if err != nil {
		return err
	}

	newAccount = makeAccount(user)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			jsonErr := WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
			if jsonErr != nil {
				fmt.Println("Error : ", jsonErr.Error())
			}
		}
	}
}

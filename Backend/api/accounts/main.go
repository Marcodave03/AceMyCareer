package accounts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KRook0110/PSD/rentaltool/backend/api/utils"
	"github.com/gorilla/mux"
)


func HandleAccounts(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	switch r.Method {
    case http.MethodGet:
        // handleGetAccount(w, r, db)
        handleGetAllAccounts(w,r,db)
        break
    case http.MethodPost:
        handleMakeAccount(w,r,db)
        break
    case http.MethodPatch:
        handleUpdateAccount(w,r,db);
        break
    case http.MethodDelete:
        handleDeleteAccount(w,r, db)
        break
    }

}

func HandleGetAccountbyUsername(w http.ResponseWriter, r *http.Request, db* sql.DB) {
    requestedUsername := mux.Vars(r)["username"]
    newAccount, err:= GetAccountByUsername(db, requestedUsername)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Println("Checkpoint 1")
    fmt.Println(newAccount)
    if err := utils.WriteJson(w, http.StatusOK, newAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    fmt.Println("Added Account", *newAccount)
}

func handleMakeAccount(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var newAccountInfo MakeAccountRequest
    err := json.NewDecoder(r.Body).Decode(&newAccountInfo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := CreateAccount(db, newAccountInfo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func handleUpdateAccount(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var request MakeAccountRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    authorized, err := ValidateAccount(db, ValidateAccountRequest{
        Username: request.Username,
        PasswordToken: request.PasswordToken,
        PasswordSalt: request.PasswordSalt,
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if !authorized {
        http.Error(w, "Not Authorized", http.StatusUnauthorized)
        return
    }
    if err := UpdateAccount(db, request); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func handleDeleteAccount(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var credentialInfo ValidateAccountRequest
    if err := json.NewDecoder(r.Body).Decode(&credentialInfo); err != nil {
        fmt.Println("Error bad request")
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    authorized, err := ValidateAccount(db, credentialInfo)
    if err != nil {
        fmt.Println("Error while checking account")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if !authorized {
        fmt.Println("Unauthorized")
        http.Error(w, "Not Authorized", http.StatusUnauthorized)
        return
    }

    if err := DeleteAccountWithUsername(db, credentialInfo.Username); err != nil {
        fmt.Println("Error while deleting account")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func handleGetAllAccounts(w http.ResponseWriter, r *http.Request, db *sql.DB) {

    accounts, err := getAllAccounts(db)
    fmt.Println("accounts total (handleGetAllAccounts) : ", len(accounts))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return;
    }

    err = utils.WriteJson(w, http.StatusOK, accounts)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return;
    }
}

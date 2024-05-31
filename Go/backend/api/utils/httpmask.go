package utils

import (
	"database/sql"
	"net/http"
)

type httpHandlerFuncMask func(http.ResponseWriter, *http.Request)
type apiFuncMask func(http.ResponseWriter, *http.Request, *sql.DB)

func MakeHttpHandlerFunc(f apiFuncMask,db *sql.DB ) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        f(w,r,db)
    }
}


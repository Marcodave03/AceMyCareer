package products

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KRook0110/PSD/rentaltool/backend/api/utils"
	_ "github.com/lib/pq"
)


func HandleProducts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
    case http.MethodGet:
        HandleGetAllProducts(w,r,db)
        break
    case http.MethodPost:
        HandlePostProduct(w,r,db)
        break
    case http.MethodPatch:
        break
    case http.MethodDelete:
        HandleDeleteProduct(w, r, db)
        break
    }
}

func HandleProductTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    switch r.Method {
    case http.MethodGet:
        HandleGetAllProducts(w,r,db)
        break
    case http.MethodPost:
        HandlePostProductTypes(w,r,db)
        break
    case http.MethodPatch:
        break
    case http.MethodDelete:
        break
}
}

func HandlePostProductTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var requestProductType ProductType
    if err := json.NewDecoder(r.Body).Decode(&requestProductType); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := CreateProductType(db, requestProductType); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.WriteJson(w, http.StatusOK, "Product Type Added")
}

func HandleGetProductTypes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    productTypes, err := GetAllProductTypes(db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err :=  utils.WriteJson(w, http.StatusOK, productTypes); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func HandleDeleteProduct(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var request DeleteProductRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w,  err.Error(), http.StatusBadRequest)
        return
    }
    if err := DeleteProduct(db, request); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)

}

func HandlePostProduct(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var product Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := CreateProduct(db, product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("Created product = ", product)

    w.WriteHeader(http.StatusOK)
}

func HandleGetAllProducts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    products, err := GetAllProducts(db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := utils.WriteJson(w, http.StatusOK, products) ; err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


package api

import (
	"log"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"github.com/KRook0110/PSD/rentaltool/backend/api/accounts"
	"github.com/KRook0110/PSD/rentaltool/backend/api/utils"
	"github.com/KRook0110/PSD/rentaltool/backend/api/products"
	"github.com/KRook0110/PSD/rentaltool/backend/api/carts"

)

type ApiServer struct {
    listenPort string
    storage *sql.DB
}


func (s* ApiServer) handleCreateAll(w http.ResponseWriter, r *http.Request ) {
    accounts.CreateTableAccounts(s.storage)
    products.CreateTableProductCategories(s.storage)
    products.CreateTableProductTypes(s.storage)
    products.CreateTableProducts(s.storage)
    carts.CreateTableCarts(s.storage)
}


func (s* ApiServer) handleDropAll(w http.ResponseWriter, r *http.Request ) {
    carts.DropTableCarts(s.storage)
    products.DropTableProducts(s.storage)
    products.DropTableProductTypes(s.storage)
}


func CreateApiServer(listenPort string) *ApiServer{
    connStr := "user=postgres dbname=postgres password='#rEnt172635' sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err.Error())
    }

    if err != nil {
        log.Fatal(err.Error())
    }

    return &ApiServer{
        listenPort: listenPort,
        storage: db,
    }
}

func (s *ApiServer) Run() {

    router := mux.NewRouter();

    router.HandleFunc("/api/accounts", utils.MakeHttpHandlerFunc(accounts.HandleAccounts, s.storage))
    router.HandleFunc("/api/accounts/{username}", utils.MakeHttpHandlerFunc(accounts.HandleGetAccountbyUsername, s.storage))

    // products
    router.HandleFunc("/api/products", utils.MakeHttpHandlerFunc(products.HandleProducts,s.storage))
    router.HandleFunc(("/api/producttypes", util.MakeHttpHandlerFunc()))

    router.HandleFunc("/api/dropall", s.handleDropAll)
    router.HandleFunc("/api/createall", s.handleCreateAll)

    http.ListenAndServe(s.listenPort, router)

    // router.HandleFunc()
}

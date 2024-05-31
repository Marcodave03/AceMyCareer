package carts

import (
	"database/sql"

	_ "github.com/lib/pq"
    "log"
)

func CreateTableCarts(db *sql.DB) {
    query := `
        CREATE SCHEMA IF NOT EXISTS Users;
        CREATE TABLE IF NOT EXISTS Users.Carts (
            cart_id SERIAL PRIMARY KEY,
            username VARCHAR(30) REFERENCES Users.Accounts(username) ON DELETE CASCADE NOT NULL,
            product_id INT REFERENCES Products.Products(product_id) ON DELETE CASCADE NOT NULL
        );
        CREATE INDEX IF NOT EXISTS carts_cartid_idx ON Users.Carts (username);
    `
    _, err := db.Exec(query);
    if err != nil {
        log.Fatal(err.Error());
    }
}

func DropTableCarts(db *sql.DB) {
    query := `DROP TABLE IF EXISTS Users.Carts;`
    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err.Error())
    }
}

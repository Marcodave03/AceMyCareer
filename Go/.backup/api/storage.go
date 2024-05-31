package api

import (
	"database/sql"

	"crypto/sha512"
	"encoding/hex"
	_ "github.com/lib/pq"
)

const NameCharLength = 20

type Storage interface {
	CreateAccount( *Account ) error
	DelelteAccount( string ) error
	UpdateAccount( *Account ) error
	GetAccountByUsername( string ) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error){
	connStr := "user=postgres dbname=postgres password=#rEnt172635 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func ( s *PostgressStore ) createAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS account (
			username VARCHAR($1) PRIMARY KEY,
			first_name VARCHAR($1) NOT NULL,
			last_name VARCHAR($1),
			password_token BYTEA NOT NULL,
			password_salt BYTEA NOT NULL
		);
	`
	_, err := s.db.Exec(query, NameCharLength);
	return err
}

func ( s *PostgressStore ) CreateAccount( *Account ) error {
	return nil
}
func ( s *PostgressStore ) 	DelelteAccount( string ) error{
	return nil
}
func ( s *PostgressStore ) UpdateAccount( *Account ) error {
	return nil
}
func ( s* PostgressStore ) GetAccountByUsername( string ) (*Account, error) {
	return nil, nil
}






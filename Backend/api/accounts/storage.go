package accounts

import (
	"database/sql"
	"errors"

	"log"

	_ "github.com/lib/pq"
)

func CreateTableAccounts(db *sql.DB) { // add not null
	query := `
        CREATE SCHEMA IF NOT EXISTS Users;
        CREATE TABLE IF NOT EXISTS Users.Accounts (
            username VARCHAR(30) PRIMARY KEY,
            firstname VARCHAR(30),
            lastname VARCHAR(30),
            password_token BYTEA,
            password_salt CHAR(5),
            profile_picture_url VARCHAR(50)
        );
    `
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err.Error())
	}
}

func getAllAccounts(db *sql.DB) ([]Account, error) {
	query := `
        SELECT * FROM Users.Accounts;
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var account Account
		err := rows.Scan(
			&account.Username,
			&account.FirstName,
			&account.LastName,
			&account.PasswordToken,
			&account.PasswordSalt,
			&account.ProfilePictureUrl,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
func ValidateAccount(db *sql.DB, reqInfo ValidateAccountRequest) ( bool,error ) {
    query := `SELECT Users.Accounts.username FROM Users.Accounts WHERE Users.Accounts.username = $1 AND Users.Accounts.password_token = $2;`
    rows, err := db.Query(query, reqInfo.Username, reqInfo.PasswordToken);
    if err != nil {
        return false, err
    }
    defer rows.Close()

    return rows.Next(), rows.Err()
}


func CreateAccount(db *sql.DB, account MakeAccountRequest) error {
	query := `
        INSERT INTO Users.Accounts (username, firstname, lastname, password_token, password_salt, profile_picture_url)
        VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6
        );
    `
	_, err := db.Exec(
		query,
		account.Username,
		account.FirstName,
		account.LastName,
		account.PasswordToken,
		account.PasswordSalt,
		account.ProfilePictureUrl,
	)

	return err
}

func DeleteAccountWithUsername(db *sql.DB, username string) error {
    query := `
    DELETE FROM Users.Accounts WHERE Users.Accounts.username = $1;
    `
    _, err := db.Exec(query, username)
    return err
}

func GetAccountByUsername(db *sql.DB, username string) (*Account, error) {
	query := `
    SELECT * FROM Users.Accounts WHERE username = $1;
    `
    account := new(Account)
	rows, err := db.Query(query, username)
	if err != nil {
		return nil, err
	}
    defer rows.Close()

    if rows.Next() == false{
        return nil, errors.New("User Not Found")
    }
    err = rows.Scan(
        &account.Username,
        &account.FirstName,
        &account.LastName,
        &account.PasswordToken,
        &account.PasswordSalt,
        &account.ProfilePictureUrl,
        )
    if err != nil {
        return nil, err
    }

	return account, nil
}

func UpdateAccount(db *sql.DB, newUserData MakeAccountRequest) error {
    query := `
    UPDATE Users.Accounts
    SET firstname = $1,
        lastname = $2,
        profile_picture_url = $3
    WHERE username = $4;
    `

    _, err := db.Exec(query, newUserData.FirstName, newUserData.LastName, newUserData.ProfilePictureUrl, newUserData.Username)
    return err
}

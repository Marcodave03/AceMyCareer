/*
   TODO:
   - Best practice on user passwords(SHA(frontend), AES, RSA)
*/

package users

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// gw ga make this secure karena tidak cukup waktu
func CreateTableUsers(db *sql.DB) error {
	query := `
    CREATE SCHEMA IF NOT EXISTS users;
    CREATE TABLE IF NOT EXISTS users.users (
        username VARCHAR(30) PRIMARY KEY,
        password VARCHAR(30),
        firstname VARCHAR(40),
        lastname VARCHAR(40),
        profile_picture_url VARCHAR(50) DEFAULT "/images/profile/default_profile_picture.png",
        score_1 INT,
        score_2 INT,
        score_3 INT,
        score_4 INT
    );
    `
	_, err := db.Exec(query)
	return err
}

// TODO:
func InsertTableUsers(db *sql.DB) error {
    // query := `
    // INSERT INTO users.users(username)
    // `
    return nil
}

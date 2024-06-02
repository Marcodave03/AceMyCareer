/*
   TODO:
   - Best practice on user passwords(SHA(frontend), AES, RSA)
*/

package users

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// gw ga make this secure karena tidak cukup waktu
func CreateTableUsers(db *sql.DB) error {
	query := `
    CREATE SCHEMA IF NOT EXISTS accounts;
    CREATE TABLE IF NOT EXISTS accounts.users (
        username VARCHAR(30) PRIMARY KEY,
        password VARCHAR(30) NOT NULL,
        firstname VARCHAR(30),
        lastname VARCHAR(30),
        email VARCHAR(30),
        profile_picture_url VARCHAR(50),
        score_technical INT DEFAULT 0,
        score_leadership INT DEFAULT 0,
        score_teamwork INT DEFAULT 0,
        score_organization INT DEFAULT 0
    );
    `
	_, err := db.Exec(query)
	return err
}

// TODO:
func insertUser(db *sql.DB, insertedUserData User) error {
	query := `
    INSERT INTO accounts.users (
    username,
    password,
    firstname,
    lastname,
    email,
    profile_picture_url,
    score_technical,
    score_leadership,
    score_teamwork,
    score_organization
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `
	_, err := db.Exec(query,
		insertedUserData.Username,
		insertedUserData.Password,
		insertedUserData.Firstname,
		insertedUserData.Lastname,
		insertedUserData.Email,
		insertedUserData.ProfilePictureUrl,
		insertedUserData.ScoreTechnical,
		insertedUserData.ScoreLeadership,
		insertedUserData.ScoreTeamwork,
		insertedUserData.ScoreOrganization,
	)

	return err
}

// if already in table then true
func checkUserUsernameInUserTable(db *sql.DB, checkUsername string) (bool, error) {
	var queriedUsername string
	err := db.QueryRow("SELECT username FROM accounts.users WHERE username = $1;", checkUsername).Scan(&queriedUsername)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if queriedUsername != checkUsername {
		fmt.Println("Anomaly on CheckTableUserForUsername: queriedUsername is not the same as checkUsername")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func checkUserCredentials(db *sql.DB, usercredentials UserCredentialRequest) (bool, error) {
    username := usercredentials.Username
    password := usercredentials.Password
	var queriedUsername string
	err := db.QueryRow("SELECT username FROM accounts.users WHERE username = $1 AND password = $2;", username, password).Scan(&queriedUsername)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if queriedUsername != username {
		fmt.Println("Anomaly on CheckTableUserForUsername: queriedUsername is not the same as checkUsername")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func deleteUser(db *sql.DB, targetUsername string) error {
	query := `
    DELETE FROM accounts.users WHERE username = $1;
    `
	_, err := db.Exec(query, targetUsername)
	return err
}

func getAllUsersFromTableUser(db *sql.DB) ([]User, error) {
	query := `
    SELECT * FROM accounts.users;
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Username,
			&user.Password,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
			&user.ProfilePictureUrl,
			&user.ScoreTechnical,
			&user.ScoreLeadership,
			&user.ScoreTeamwork,
			&user.ScoreOrganization)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

    return users, nil
}

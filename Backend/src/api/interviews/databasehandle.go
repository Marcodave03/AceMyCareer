package interviews

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateTableInterviews(db *sql.DB) error {
	query := `
    CREATE SCHEMA IF NOT EXISTS interviews;

    CREATE TABLE IF NOT EXISTS interviews.tags (
        name VARCHAR(30) PRIMARY KEY
    );


    CREATE TABLE IF NOT EXISTS interviews.interview_levels (
        id SERIAL PRIMARY KEY,
        name VARCHAR(30) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS interviews.interview_industries (
        name VARCHAR(30) PRIMARY KEY
    );

    CREATE TABLE IF NOT EXISTS interviews.interview_positions (
        name VARCHAR(30) PRIMARY KEY
    );

    CREATE TABLE IF NOT EXISTS interviews.technical_requirements (
        name VARCHAR(30) PRIMARY KEY
    );


    CREATE TABLE IF NOT EXISTS interviews.interviews (
        id SERIAL PRIMARY KEY,
        title VARCHAR(30),
        image_url VARCHAR(50),
        industry_name VARCHAR(30) REFERENCES interviews.interview_industries(name),
        position_name VARCHAR(30) REFERENCES interviews.interview_positions(name),
        level_id SERIAL REFERENCES interviews.interview_levels(id),
        experience TEXT,
        likes INT DEFAULT 0
    );

    CREATE TABLE IF NOT EXISTS interviews.interview_tags (
        id SERIAL PRIMARY KEY,
        interview_id SERIAL REFERENCES interviews.interviews(id),
        tag_name VARCHAR(30) REFERENCES interviews.tags(name)
    );

    CREATE TABLE IF NOT EXISTS interviews.interview_requirements (
        id SERIAL PRIMARY KEY,
        requirement_name VARCHAR(30) REFERENCES interviews.technical_requirements,
        interview_id SERIAL REFERENCES interviews.interviews(id)
    );
    `
	_, err := db.Exec(query)
	return err
}

func getAllInterviewLevels(db *sql.DB) ( []interviewLevel,  error ) {
    query := `
    SELECT * FROM interviews.interview_levels;
    `
    rows, err := db.Query(query);
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var levels []interviewLevel
    for rows.Next() {
        var curLevel interviewLevel
        err := rows.Scan(&curLevel.ID, &curLevel.Name)
        if err != nil {
            return nil, err
        }
        levels = append(levels, curLevel)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return levels, nil
}

// inserts level name
func insertLevel(db *sql.DB, newLevelName string) error {
    query :=` INSERT INTO interviews.interview_levels (name) VALUES ($1);`

    _, err := db.Exec(query, newLevelName)
    return err
}

func deleteLevel(db *sql.DB, levelID int) error {
    query := `SELECT id FROM interviews.interview_levels WHERE id = $1;`
    var foundid int
    if err := db.QueryRow(query, levelID).Scan(&foundid); err != nil {
        return err
    }

    query = `DELETE FROM interviews.interview_levels WHERE id = $1`
    _, err := db.Exec(query, levelID)
    return err
}

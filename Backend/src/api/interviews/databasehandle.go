package interviews

import (
	"database/sql"
	"fmt"

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

// Level {{{
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

func deleteLevel(db *sql.DB, levelID int) error {
    query := `DELETE FROM interviews.interview_levels WHERE id = $1`
    affected, err := db.Exec(query, levelID)
    if err != nil {
        return err
    }
    total, err := affected.RowsAffected()
    if total == 0 {
        return fmt.Errorf("Not Found")
    }

    return err
}


// inserts level name
func insertLevel(db *sql.DB, newLevelName string) error {
    query :=` INSERT INTO interviews.interview_levels (name) VALUES ($1);`

    _, err := db.Exec(query, newLevelName)
    return err
}

// }}}

// Positions {{{
func getAllInterviewPositions(db *sql.DB) ( []interviewPosition,  error ) {
    query := `
    SELECT * FROM interviews.interview_positions;
    `
    rows, err := db.Query(query);
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var positions []interviewPosition
    for rows.Next() {
        var curPosition interviewPosition
        err := rows.Scan(&curPosition.Name)
        if err != nil {
            return nil, err
        }

        positions = append(positions, curPosition)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return positions, nil
}

func insertPosition(db *sql.DB, newPositionName string) error {
    query :=` INSERT INTO interviews.interview_positions (name) VALUES ($1);`
    _, err := db.Exec(query, newPositionName)
    return err
}

func deletePositions(db *sql.DB, positionName string) error {
    query := `DELETE FROM interviews.interview_positions WHERE name = $1`
    result, err := db.Exec(query, positionName)
    affected, err := result.RowsAffected();
    if  affected == 0 {
        return fmt.Errorf("Not Found")
    }
    return err
}

// }}}

// Industries
func getAllInterviewIndustries(db *sql.DB) ( []interviewIndustry,  error ) {
    query := `
    SELECT * FROM interviews.interview_industries;
    `
    rows, err := db.Query(query);
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var industries []interviewIndustry
    for rows.Next() {
        var curIndustrie interviewIndustry
        err := rows.Scan(&curIndustrie.Name)
        if err != nil {
            return nil, err
        }

        industries = append(industries, curIndustrie)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return industries, nil
}

func insertIndustry(db *sql.DB, newIndustryName string) error {
    query :=` INSERT INTO interviews.interview_industries (name) VALUES ($1);`
    _, err := db.Exec(query, newIndustryName)
    return err
}

func deleteIndustries(db *sql.DB, industryName string) error {
    query := `DELETE FROM interviews.interview_industries WHERE name = $1`
    affected, err := db.Exec(query, industryName)
    if err != nil {
        return err
    }
    total, err := affected.RowsAffected()
    if total == 0 {
        return fmt.Errorf("Not Found")
    }
    return err
}



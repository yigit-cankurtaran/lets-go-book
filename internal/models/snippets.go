package models

// internal is for non-specific code
// a database model can be used in any project
// so it fits the bill here

import (
	"database/sql"
	"time"
)

// Snippet is a struct that models a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// inserting snippet into database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// the SQL statement we want to execute
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	// we use UTC_TIMESTAMP() to get the current time in UTC

	// Exec() is used to execute a SQL statement that doesn't return rows
	// first parameter is the statement, the rest are the parameters
	// it returns an sql.Result object
	// sql.Result provides LastInsertId() and RowsAffected() methods
	// to get the ID of the newly inserted record and the number of rows affected
	// both are int64
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// we can completely ignore the return value, it's accepted and common
	// we just need to do _, err instead

	// LastInsertId() returns the ID of the last row inserted
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// return the ID of the new snippet
	// it's in int64 so we need to convert it to int
	return int(id), nil
}

// getting a snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// getting the 10 most recent snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

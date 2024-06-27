package models

// internal is for non-specific code
// a database model can be used in any project
// so it fits the bill here

import (
	"database/sql"
	"errors"
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
	// statement for get
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// QueryRow on the connection pool to execute the SQL statement
	// returns a pointer to a sql.Row object
	row := m.DB.QueryRow(stmt, id)
	// QueryRow is expected to return at least 1 row

	// initialize pointer to a new Snippet struct
	s := &Snippet{}

	// use row.Scan() to copy values to the Snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// if there are no rows we get a specific error
		// otherwise we'll get errors normally
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// if everything is fine, return the Snippet struct
	return s, nil
}

// getting the 10 most recent snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

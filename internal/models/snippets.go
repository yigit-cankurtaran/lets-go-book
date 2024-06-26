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
	return 0, nil
}

// getting a snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// getting the 10 most recent snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

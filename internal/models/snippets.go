package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, NOW(), NOW() + INTERVAL '1 DAY' * $3)
	RETURNING id`
	row := m.DB.QueryRow(query, title, content, expires)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > NOW() and id = $1`
	row := m.DB.QueryRow(query, id)

	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

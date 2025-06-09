/*
	Repository layer for your application

	Used by app.go
*/

package store

import "database/sql"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Published bool      `json:"published"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

type DataStore interface {
	CreatePost(*Post) (*Post, error)
	GetPostById(int64) (*Post, error)
	UpdatePost(*Post) error
	DeletePostById(int64) error
}

func (pg *PostgresStore) CreatePost(entity *Post) (*Post, error) {
	return nil, nil
}

func (pg *PostgresStore) GetPostById(id int64) (*Post, error) {
	return nil, nil
}

func (pg *PostgresStore) UpdatePost(entity *Post) error {
	return nil
}

func (pg *PostgresStore) DeletePostById(id int64) error {
	return nil
}

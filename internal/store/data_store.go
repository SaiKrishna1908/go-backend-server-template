/*
	Repository layer for your application

	Used by app.go
*/

package store

import "database/sql"

type Entity struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Entries []Entity2 `json:"entries"`
}

type Entity2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

type DataStore interface {
	CreateEntity(*Entity) (*Entity, error)
	GetEntityById(int64) (*Entity, error)
	UpdateEntity(*Entity) error
	DeleteEntityById(int64) error
}

func (pg *PostgresStore) CreateEntity(entity *Entity) (*Entity, error) {
	return nil, nil
}

func (pg *PostgresStore) GetEntityById(id int64) (*Entity, error) {
	return nil, nil
}

func (pg *PostgresStore) UpdateEntity(entity *Entity) error {
	return nil
}

func (pg *PostgresStore) DeleteEntityById(id int64) error {
	return nil
}

package store

import (
	"database/sql"
	"time"
)

type password struct {
	plainText string
	hash      []byte
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

type UserStore interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	GetUserToken(scope string, tokenPlainText string) (*User, error)
}

func (userStore *PostgresUserStore) CreateUser(user *User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := userStore.db.QueryRow(query, user.Username, user.Email, user.Password.hash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

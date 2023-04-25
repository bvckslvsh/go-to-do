package repository

import (
	"fmt"

	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/jmoiron/sqlx"
)

type AuthPostgress struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgress {
	return &AuthPostgress{db: db}
}

func (r *AuthPostgress) CreateUser(user gotodo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgress) GetUser(username, password string) (gotodo.User, error) {
	var user gotodo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err

}

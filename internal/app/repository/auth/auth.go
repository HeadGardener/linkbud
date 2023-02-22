package authrepository

import (
	"errors"
	"fmt"
	postgres "github.com/HeadGardener/linkbud/internal/app/db"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		postgres.UsersTable)

	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *AuthRepository) IfUserExist(userInput models.UserInput) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", postgres.UsersTable)

	err := r.db.Get(&id, query, userInput.Username, userInput.Password)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) IfUserExistByUN(username string) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", postgres.UsersTable)

	err := r.db.Get(&id, query, username)
	if err != nil {
		return 0, errors.New("user dont exist")
	}

	return id, nil
}

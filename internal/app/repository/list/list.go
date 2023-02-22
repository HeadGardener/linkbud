package listrepository

import (
	"errors"
	"fmt"
	postgres "github.com/HeadGardener/linkbud/internal/app/db"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type ListRepository struct {
	db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) Create(userID int, list models.LinkList) (int, error) {
	var listID int

	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1 AND list_title=$2", postgres.UsersListsTable)
	err := r.db.Get(&listID, query, userID, list.ShortTitle)
	if err == nil {
		return 0, errors.New("you already use this title")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, short_title, description) VALUES ($1, $2, $3) RETURNING id",
		postgres.ListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.ShortTitle, list.Description)
	if err := row.Scan(&listID); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_title) VALUES ($1, $2)", postgres.UsersListsTable)
	_, err = tx.Exec(usersListQuery, userID, list.ShortTitle)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listID, tx.Commit()
}

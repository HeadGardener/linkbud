package listrepository

import (
	"errors"
	"fmt"
	postgres "github.com/HeadGardener/linkbud/internal/app/db"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ListRepository struct {
	db *sqlx.DB
}

func NewListRepository(db *sqlx.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) Create(userID int, list models.List) (int, error) {
	var listID int

	checkIfTitleExistsQuery := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1 AND list_title=$2",
		postgres.UsersListsTable)

	err := r.db.Get(&listID, checkIfTitleExistsQuery, userID, list.ShortTitle)
	if err == nil || listID != 0 {
		return 0, errors.New("you already use this title")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf(`INSERT INTO %s (title, short_title, description)
											VALUES ($1, $2, $3) RETURNING id`,
		postgres.ListsTable)

	row := tx.QueryRow(createListQuery, list.Title, list.ShortTitle, list.Description)
	if err := row.Scan(&listID); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id, list_title) VALUES ($1, $2, $3)",
		postgres.UsersListsTable)

	_, err = tx.Exec(usersListQuery, userID, listID, list.ShortTitle)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listID, tx.Commit()
}

func (r *ListRepository) GetAll(userID int) ([]models.List, error) {
	var ids []int

	getIDsQuery := fmt.Sprintf("SELECT list_id FROM %s WHERE user_id=$1",
		postgres.UsersListsTable)

	if err := r.db.Select(&ids, getIDsQuery, userID); err != nil {
		return nil, err
	}

	var lists []models.List
	var list models.List

	getListByIDQuery := fmt.Sprintf("SELECT id, title, short_title, description FROM %s WHERE id=$1",
		postgres.ListsTable)

	for _, n := range ids {
		err := r.db.Get(&list, getListByIDQuery, n)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *ListRepository) GetList(userID int, title string) (models.List, error) {
	var listID int

	getIDQuery := fmt.Sprintf("SELECT list_id FROM %s WHERE user_id=$1 AND list_title=$2",
		postgres.UsersListsTable)

	if err := r.db.Get(&listID, getIDQuery, userID, title); err != nil {
		return models.List{}, err
	}

	var list models.List

	getListByIDQuery := fmt.Sprintf("SELECT id, title, short_title, description FROM %s WHERE id=$1",
		postgres.ListsTable)

	err := r.db.Get(&list, getListByIDQuery, listID)

	return list, err
}

func (r *ListRepository) Update(userID int, title string, list models.ListInput) (int, error) {
	var listID int

	updateUsersListQuery := fmt.Sprintf(`UPDATE %s SET list_title=$1 
												WHERE user_id=$2 AND list_title=$3 RETURNING list_id`,
		postgres.UsersListsTable)

	err := r.db.QueryRow(updateUsersListQuery, list.ShortTitle, userID, title).Scan(&listID)
	if err != nil {
		return 0, err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if list.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, list.Title)
		argID++
		setValues = append(setValues, fmt.Sprintf("short_title=$%d", argID))
		args = append(args, list.ShortTitle)
		argID++
	}

	if list.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, list.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s b SET %s WHERE b.id = $%d",
		postgres.ListsTable, setQuery, argID)
	args = append(args, listID)

	_, err = r.db.Exec(query, args...)

	return listID, err
}

func (r *ListRepository) Delete(userID int, title string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listID int

	deleteFromUsersListsQuery := fmt.Sprintf(`DELETE FROM %s 
													WHERE user_id=$1 AND list_title=$2 RETURNING list_id`,
		postgres.UsersListsTable)

	err = tx.QueryRow(deleteFromUsersListsQuery, userID, title).Scan(&listID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	deleteFromListsQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 AND short_title=$2`,
		postgres.ListsTable)
	_, err = tx.Exec(deleteFromListsQuery, listID, title)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listID, tx.Commit()
}

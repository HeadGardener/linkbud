package linkrepository

import (
	"errors"
	"fmt"
	postgres "github.com/HeadGardener/linkbud/internal/app/db"
	"github.com/HeadGardener/linkbud/internal/app/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type LinkRepository struct {
	db *sqlx.DB
}

func NewLinkRepository(db *sqlx.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(link models.Link, listID int) (int, error) {
	var id int

	checkIfTitleExistsQuery := fmt.Sprintf("SELECT id FROM %s WHERE list_id=$1 AND link_title=$2",
		postgres.ListsLinksTable)

	err := r.db.Get(&id, checkIfTitleExistsQuery, listID, link.ShortTitle)
	if err == nil || id != 0 {
		return 0, errors.New("you already use this title for link")
	}

	var linkID int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createLinkQuery := fmt.Sprintf(`INSERT INTO %s (title, short_title, url)
											VALUES ($1, $2, $3) RETURNING id`,
		postgres.LinksTable)

	row := tx.QueryRow(createLinkQuery, link.Title, link.ShortTitle, link.URL)
	if err := row.Scan(&linkID); err != nil {
		tx.Rollback()
		return 0, err
	}

	listsLinksQuery := fmt.Sprintf("INSERT INTO %s (list_id, link_id, link_title) VALUES ($1, $2, $3)",
		postgres.ListsLinksTable)

	_, err = tx.Exec(listsLinksQuery, listID, linkID, link.ShortTitle)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return linkID, tx.Commit()
}

func (r *LinkRepository) GetAll(listID int) ([]models.Link, error) {
	var links []models.Link
	query := fmt.Sprintf(`SELECT l.id, l.title, l.short_title, l.url FROM %s
								l INNER JOIN %s ll on l.id = ll.link_id WHERE ll.list_id=$1`,
		postgres.LinksTable, postgres.ListsLinksTable)

	err := r.db.Select(&links, query, listID)

	return links, err
}

func (r *LinkRepository) GetByID(listID int, linkTitle string) (models.Link, error) {
	var linkID int

	getIDQuery := fmt.Sprintf("SELECT link_id FROM %s WHERE list_id=$1 AND link_title=$2",
		postgres.ListsLinksTable)

	if err := r.db.Get(&linkID, getIDQuery, listID, linkTitle); err != nil {
		return models.Link{}, err
	}

	var link models.Link

	getListByIDQuery := fmt.Sprintf("SELECT id, title, short_title, url FROM %s WHERE id=$1",
		postgres.LinksTable)

	err := r.db.Get(&link, getListByIDQuery, linkID)

	return link, err
}

func (r *LinkRepository) Update(listID int, linkInput models.LinkInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if linkInput.ShortTitle != nil {
		updateUsersListQuery := fmt.Sprintf(`UPDATE %s SET link_title=$1 
												WHERE list_id=$2 AND link_id=$3`,
			postgres.ListsLinksTable)

		_, err = tx.Exec(updateUsersListQuery, linkInput.ShortTitle, listID, linkInput.ID)
		if err != nil {
			return err
		}
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if linkInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, linkInput.Title)
		argID++
	}

	if linkInput.ShortTitle != nil {
		setValues = append(setValues, fmt.Sprintf("short_title=$%d", argID))
		args = append(args, linkInput.ShortTitle)
		argID++
	}

	if linkInput.URL != nil {
		setValues = append(setValues, fmt.Sprintf("url=$%d", argID))
		args = append(args, linkInput.URL)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s l SET %s WHERE l.id = $%d",
		postgres.LinksTable, setQuery, argID)
	args = append(args, linkInput.ID)

	_, err = tx.Exec(query, args...)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *LinkRepository) Delete(linkID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteFromUsersListsQuery := fmt.Sprintf(`DELETE FROM %s 
													WHERE link_id=$1`,
		postgres.ListsLinksTable)

	_, err = tx.Exec(deleteFromUsersListsQuery, linkID)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteFromListsQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`,
		postgres.LinksTable)
	_, err = tx.Exec(deleteFromListsQuery, linkID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

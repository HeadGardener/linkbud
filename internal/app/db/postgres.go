package postgres

import (
	"fmt"
	"github.com/HeadGardener/linkbud/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	UsersTable      = "users"
	ListsTable      = "lists"
	UsersListsTable = "users_lists"
	LinksTable      = "links"
	ListsLinksTable = "lists_links"
)

func NewDB(conf configs.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s dbname=%s sslmode=%s", conf.Host, conf.DBName, conf.SSLMode))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

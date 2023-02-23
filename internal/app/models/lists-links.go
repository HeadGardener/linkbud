package models

type List struct {
	ID          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	ShortTitle  string `json:"short_title" db:"short_title"`
	Description string `json:"description" db:"description"`
}

type ListInput struct {
	Title       *string `json:"title" db:"title" binding:"required"`
	ShortTitle  string  `json:"short_title" db:"short_title"`
	Description *string `json:"description" db:"description"`
}

type UsersList struct {
	ID     int
	UserID int
	ListID int
}

type Link struct {
	ID    int    `json:"-" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	URL   string `json:"url" db:"url"`
}

type ListsLinks struct {
	ID     int
	ListID int
	LinkID int
}

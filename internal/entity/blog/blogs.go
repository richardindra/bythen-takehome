package blog

import "time"

type Blog struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	AuthorID  int64     `db:"author_id" json:"author_id"`
	Status    string    `db:"status" json:"status"`
	ViewCount int       `db:"view_count" json:"view_count"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Username  string    `db:"username" json:"username"`
	Name      string    `db:"name" json:"name"`
}

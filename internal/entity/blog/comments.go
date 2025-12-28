package blog

import "time"

type Comments struct {
	ID        int64     `db:"id" json:"id"`
	BlogID    int64     `db:"blog_id" json:"blog_id"`
	AuthorID  int64     `db:"author_id" json:"author_id"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Username  string    `db:"username" json:"username"`
	Name      string    `db:"name" json:"name"`
}

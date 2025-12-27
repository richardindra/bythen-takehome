package blog

import "time"

type User struct {
	ID          int64     `db:"id" json:"id"`
	Username    string    `db:"username" json:"username"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password_hash" json:"password"`
	Status      string    `db:"status" json:"status"`
	LastLoginAt time.Time `db:"last_login_at" json:"last_login_at"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type RespCreateUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

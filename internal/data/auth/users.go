package auth

import (
	"bythen-takehome/internal/entity/blog"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func (d Data) CreateUser(ctx context.Context, user blog.User) (int64, error) {
	var (
		userID int64
	)
	res, err := (*d.stmt)[createUser].ExecContext(ctx,
		user.Username,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return 0, fmt.Errorf("[DATA][CreateUser]: %w", err)

	}
	userID, _ = res.LastInsertId()

	return userID, nil
}

func (d Data) CheckUser(ctx context.Context, username, email string) (int, error) {
	var (
		count int
		err   error
	)

	err = (*d.stmt)[checkUser].QueryRowxContext(ctx, username, email).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("[DATA][CheckUser]: %w", err)
	}

	return count, nil
}

func (d Data) GetUserByUsername(ctx context.Context, username string) (blog.User, error) {
	user := blog.User{}

	if err := (*d.stmt)[getUserByUsername].QueryRowxContext(ctx, username).StructScan(&user); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("[DATA][GetUserByUsername]: %w", errors.New("Username not found"))
		}

		return user, fmt.Errorf("[DATA][GetUserByUsername]: %w", err)
	}

	return user, nil
}

func (d Data) UpdateLastLogin(ctx context.Context, username string) (time.Time, error) {
	var lastLoginAt time.Time

	_, err := (*d.stmt)[updateLastLogin].ExecContext(ctx, username)
	if err != nil {
		return time.Time{}, fmt.Errorf("[DATA][UpdateLastLogin]: %w", err)
	}

	err = (*d.stmt)[getLastLogin].QueryRowxContext(ctx, username).Scan(&lastLoginAt)
	if err != nil {
		return time.Time{}, fmt.Errorf("[DATA][UpdateLastLogin]: %w", err)
	}

	return lastLoginAt, nil
}

package blog

import (
	"bythen-takehome/internal/entity/blog"
	"bythen-takehome/pkg/errors"
	"context"
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
		return 0, errors.Wrap(err, "[DATA][CreateUser]")
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
		return count, errors.Wrap(err, "[DATA][CheckUser]")
	}

	return count, nil
}

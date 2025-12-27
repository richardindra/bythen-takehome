package blog

import (
	"bythen-takehome/internal/entity/blog"
	"bythen-takehome/pkg/errors"
	"context"

	"github.com/raja/argon2pw"
)

func (s Service) CreateUser(ctx context.Context, req blog.User) (blog.RespCreateUser, error) {
	var (
		resp blog.RespCreateUser
	)
	if req.Email == "" {
		return resp, errors.Wrap(errors.New("email is empty"), "[SERVICE][CreateUser]")
	} else if req.Password == "" {
		return resp, errors.Wrap(errors.New("password is empty"), "[SERVICE][CreateUser]")
	}

	checkUser, err := s.data.CheckUser(ctx, req.Username, req.Email)
	if err != nil {
		return resp, errors.Wrap(err, "[SERVICE][CreateUser]")
	}

	if checkUser > 0 {
		return resp, errors.Wrap(errors.New("username/email has already been used"), "[SERVICE][CreateUser]")
	} else {
		hashedPassword, err := argon2pw.GenerateSaltedHash(req.Password)
		if err != nil {
			return resp, errors.Wrap(err, "[SERVICE][CreateUser][HASHING]")
		}

		user := blog.User{
			Username: req.Username,
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
		}
		lastID, err := s.data.CreateUser(ctx, user)
		if err != nil {
			return resp, errors.Wrap(err, "[SERVICE][CreateUser]")
		}

		resp = blog.RespCreateUser{
			ID:       lastID,
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		}

		return resp, nil
	}
}

package auth

import (
	"bythen-takehome/internal/entity/blog"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Data interface {
	// User
	CreateUser(ctx context.Context, user blog.User) (int64, error)
	CheckUser(ctx context.Context, username, email string) (int, error)
	GetUserByUsername(ctx context.Context, username string) (blog.User, error)
	UpdateLastLogin(ctx context.Context, username string) (time.Time, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}

func GenerateToken(userid int64, username, name string, exp time.Time) (string, error) {
	var secretKey = []byte("secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid":   userid,
			"username": username,
			"name":     name,
			"exp":      exp.Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

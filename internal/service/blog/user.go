package blog

import (
	"bythen-takehome/internal/entity/auth"
	"bythen-takehome/internal/entity/blog"
	"bythen-takehome/pkg/errors"
	"context"
	"strings"
	"time"

	"github.com/raja/argon2pw"
)

func (s Service) Register(ctx context.Context, req blog.User) (blog.RespCreateUser, error) {
	var (
		resp blog.RespCreateUser
	)
	if req.Email == "" {
		return resp, errors.New("Email is required")
	} else if req.Password == "" {
		return resp, errors.New("Passsword is required")
	}

	checkUser, err := s.data.CheckUser(ctx, req.Username, req.Email)
	if err != nil {
		return resp, errors.Wrap(err, "[SERVICE][CreateUser]")
	}

	if checkUser > 0 {
		return resp, errors.New("Username/email has already been used")
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

func isValidArgon2Hash(hash string) bool {
	return strings.Contains(hash, "argon2")
}

func (s Service) Login(ctx context.Context, req blog.LoginRequest) (auth.LoginResponse, blog.UserInfo, error) {
	var (
		data     auth.LoginResponse
		metadata blog.UserInfo
	)

	if req.Username == "" || req.Password == "" {
		return data, metadata, errors.New("Username and password are required")
	}

	userCreds, err := s.data.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return data, metadata, errors.Wrap(err, "[SERVICE][Login]")
	}

	if userCreds.ID == 0 {
		return data, metadata, errors.New("Invalid credentials")
	} else if userCreds.Status == "N" {
		return data, metadata, errors.New("User is inactive")
	}

	//optional utk cek password argon atau bukan biar ga panic
	if !isValidArgon2Hash(userCreds.Password) {
		return data, metadata, errors.New("Invalid argon credentials")
	}

	valid, err := argon2pw.CompareHashWithPassword(userCreds.Password, req.Password)
	if err != nil {
		return data, metadata, err
	}

	if valid {
		t := time.Now()
		d := 12 * time.Hour
		e := t.Add(d)

		_ = s.data.UpdateLastLogin(ctx, req.Username)
		tokenString, err := GenerateToken(userCreds.ID, userCreds.Username, userCreds.Name, e)
		if err != nil {
			return data, metadata, errors.Wrap(err, "[SERVICE][Login]")
		}

		data = auth.LoginResponse{
			Message:     "Login success",
			AccessToken: tokenString,
			ExpiresAt:   e.Unix(),
		}

		metadata = blog.UserInfo{
			ID:          userCreds.ID,
			Username:    userCreds.Username,
			Name:        userCreds.Name,
			Email:       userCreds.Email,
			Status:      userCreds.Status,
			LastLoginAt: time.Now(),
			CreatedAt:   userCreds.CreatedAt,
			UpdatedAt:   userCreds.UpdatedAt,
		}

		return data, metadata, nil
	} else {
		return data, metadata, errors.New("Invalid password")
	}
}

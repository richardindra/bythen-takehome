package blog

import (
	"bythen-takehome/internal/entity/auth"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userid int, username, role, fullname string) (string, error) {
	var secretKey = []byte("secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid":   userid,
			"username": username,
			"role":     role,
			"fullname": fullname,
			"exp":      time.Now().Add(time.Hour * 2).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func CheckPasswordHash(inputPassword, hashPassword string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))
// 	return err == nil
// }

func (s Service) DecodeJWT(ctx context.Context, authToken string) (auth.Claims, error) {
	var (
		data auth.Claims
	)

	decodeToken, _ := jwt.ParseWithClaims(authToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("tokenString"), nil
	})

	fmt.Println(decodeToken.Claims)

	respDecode := decodeToken.Claims.(jwt.MapClaims)

	data.UserID = int(respDecode["userid"].(float64))
	data.Username = fmt.Sprintf("%v", respDecode["username"])
	data.Role = (respDecode["role"].(string))
	exp := int64(respDecode["exp"].(float64))

	fmt.Println(exp)

	return data, nil
}

func (s Service) ValidateTalentRole(role string) bool {
	return role == "TALENTS"
}

func (s Service) ValidateEmployerRole(role string) bool {
	return role == "EMPLOYERS"
}

func (s Service) GetJWTDetail(authToken string) (auth.DecodeJWT, error) {
	var (
		data auth.DecodeJWT
	)

	decodeToken, _ := jwt.ParseWithClaims(authToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("tokenString"), nil
	})

	respDecode := decodeToken.Claims.(jwt.MapClaims)

	data.UserID = int64(respDecode["userid"].(float64))
	data.Username = (respDecode["username"].(string))
	data.ExpireIn = int64(respDecode["exp"].(float64))

	return data, nil
}

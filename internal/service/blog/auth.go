package blog

import (
	"bythen-takehome/internal/entity/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
	data.Name = (respDecode["name"].(string))
	data.ExpireIn = int64(respDecode["exp"].(float64))

	return data, nil
}

package auth

type Claims struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type DecodeJWT struct {
	UserID   int64  `json:"userid"`
	Username string `json:"username"`
	ExpireIn int64  `json:"expirein"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

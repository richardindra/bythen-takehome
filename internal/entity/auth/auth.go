package auth

type DecodeJWT struct {
	UserID   int64  `json:"userid"`
	Username string `json:"username"`
	Name     string `json:"name"`
	ExpireIn int64  `json:"expirein"`
}

type LoginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

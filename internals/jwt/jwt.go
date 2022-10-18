package jwt

import (
	"os"
	"time"

	"github.com/kataras/jwt"
)

type JWTHeader struct {
	Typ string `json:"typ"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

func CreateTestJwt(userId string) (*string, error) {
	now := time.Now()

	issuer := os.Getenv("JWT_ISSUER")
	secret := []byte(os.Getenv("JWT_SECRET"))

	standardClaims := jwt.Claims{
		Expiry:   now.Add(60 * time.Minute).Unix(),
		IssuedAt: now.Unix(),
		Issuer:   issuer,
		Subject:  userId,
	}

	header := JWTHeader{
		Typ: "JWT",
		Kid: os.Getenv("JWT_KID"),
		Alg: jwt.HS256.Name(),
	}

	token, err := jwt.SignWithHeader(jwt.HS256, secret, standardClaims, header, jwt.MaxAge(60*time.Minute))
	if err != nil {
		return nil, err
	}

	stringToken := string(token)
	return &stringToken, nil
}

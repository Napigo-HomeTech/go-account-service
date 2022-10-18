package controllers

import (
	"os"
	"time"

	"github.com/Napigo/go-account-service/internals/models"
	commonauth "github.com/Napigo/npgcommon/auth"
	commonrest "github.com/Napigo/npgcommon/rest"
	"github.com/gofiber/fiber/v2"
)

func GetJwtController(c *fiber.Ctx) error {
	c.SendStatus(200)

	resp := commonrest.RestResponse{Context: c, Payload: "Hello JWT", HTTPStatus: 200, Status: "Success"}
	return resp.SendResponse()
}

// This function is for generating JWT only meant for dev test purpose,
// please do not use this for production
func GetTestJwtController(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	now := time.Now()
	issuer := os.Getenv("JWT_ISSUER")

	secret := []byte(os.Getenv("JWT_SECRETS"))

	jwtToken := commonauth.JWTBuilder{
		Expiry:   now.Add(60 * time.Minute).Unix(),
		IssuedAt: now.Unix(),
		Issuer:   issuer,
		Subject:  userId,
		Secret:   secret,
	}

	sToken, err := jwtToken.CreateJWT()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	c.SendStatus(200)

	resp := commonrest.RestResponse{Context: c, Payload: models.SToken{Token: *sToken}, HTTPStatus: 200, Status: "Success"}
	return resp.SendResponse()
}

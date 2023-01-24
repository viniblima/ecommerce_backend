package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(id string) (string, error) {

	expireToken := time.Now().Add(time.Hour * 24).Unix()

	claims := TokenClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "bandzest-auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("PASSWORD_SECRET")))

	return tokenString, err
}

func VerifyJWT(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth != "" {

		_, err := jwt.ParseWithClaims(strings.Split(auth, "JWT ")[1], &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodECDSA)
			if !ok {
				fmt.Println("unauthorized")
			}
			return []byte(os.Getenv("PASSWORD_SECRET")), nil
		})

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"erro": err.Error(),
			})

		}

	} else {
		return c.SendStatus(401)
	}
	return c.Next()
}

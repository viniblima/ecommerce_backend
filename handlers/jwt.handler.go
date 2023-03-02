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

func GenerateJWT(id string) (map[string]interface{}, error) {

	expireToken := time.Now().Add(time.Hour * 24)

	expireRefreshToken := time.Now().Add(time.Hour * 24 * 90)

	claims := TokenClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken.Unix(),
			Issuer:    "viniblima-auth",
			Subject:   "auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("PASSWORD_SECRET")))

	//Refresh
	claimsRefresh := TokenClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireRefreshToken.Unix(),
			Issuer:    "viniblima-auth",
			Subject:   "refresh",
		},
	}
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	tokenStringRefresh, err := tokenRefresh.SignedString([]byte(os.Getenv("PASSWORD_SECRET")))

	return map[string]interface{}{
		"Token": map[string]interface{}{
			"Hash":      tokenString,
			"ExpiresIn": expireToken,
		},
		"Refresh": map[string]interface{}{
			"Hash":      tokenStringRefresh,
			"ExpiresIn": expireRefreshToken,
		},
	}, err
}

func RefreshToken(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	claims := jwt.MapClaims{}

	if auth != "" {
		split := strings.Split(auth, "JWT ")

		if len(split) < 2 {
			return c.Status(401).JSON(fiber.Map{
				"erro": "Invalid tag",
			})
		}

		parse, _ := jwt.ParseWithClaims(strings.Split(auth, "JWT ")[1], claims, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodECDSA)
			if !ok {
				fmt.Println("unauthorized")
			}
			return []byte(os.Getenv("PASSWORD_SECRET")), nil
		})

		fmt.Println(parse)

		if claims["sub"] != "refresh" {
			return c.Status(401).JSON(fiber.Map{
				"erro": "Invalid token",
			})
		}

		// if err != nil {
		// 	return c.Status(401).JSON(fiber.Map{
		// 		"erro": err.Error(),
		// 	})

		// }
	} else {
		return c.SendStatus(401)
	}

	c.Locals("userID", claims["id"])
	//return c.SendStatus(200)
	return c.Next()
}

func VerifyJWT(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	claims := jwt.MapClaims{}
	if auth != "" {
		split := strings.Split(auth, "JWT ")

		if len(split) < 2 {
			return c.Status(401).JSON(fiber.Map{
				"erro": "Invalid tag",
			})
		}

		_, err := jwt.ParseWithClaims(strings.Split(auth, "JWT ")[1], claims, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodECDSA)
			if !ok {
				fmt.Println("unauthorized")
			}
			return []byte(os.Getenv("PASSWORD_SECRET")), nil
		})

		if claims["sub"] != "auth" {
			return c.Status(401).JSON(fiber.Map{
				"erro": "Invalid token",
			})
		}

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"erro": err.Error(),
			})

		}

	} else {
		return c.SendStatus(401)
	}

	c.Locals("userID", claims["id"])
	return c.Next()
}

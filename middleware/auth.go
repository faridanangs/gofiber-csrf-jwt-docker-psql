package middleware

import (
	"fiberjwt/model"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	tokenJWT := c.Get("Authorization")
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	if tokenJWT == "" {
		// di sini saya langsung panik aja untuk nyingkat kode
		// bisa juga kirim response ke user dalam bntuk json
		return c.Status(404).SendString("token kosonggg")
	}

	claims := model.Claims{}

	// disini kita parse tokennya kemudian masukan datanya ke dalam claims
	// errornya saya skip biar cepet
	token, err := jwt.ParseWithClaims(tokenJWT, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(err.Error())
	}

	if !token.Valid {
		return c.Status(404).SendString("token is invalid")
	}

	// jika tokenyya valid baru kita next ke handler selanjutnya
	return c.Next()

}

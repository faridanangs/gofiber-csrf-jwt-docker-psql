package handler

import (
	"fiberjwt/db"
	"fiberjwt/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	db := db.Connect()
	userRequest := model.UserLogin{}
	c.BodyParser(&userRequest)

	if err := db.Debug().Model(userRequest).Where("name = ?", userRequest.Name).Error; err != nil {
		return c.Status(404).SendString("user not found")
	}

	claims := model.Claims{
		Name:     userRequest.Name,
		Password: userRequest.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return c.Status(404).SendString("signature failed")
	}

	return c.Status(200).JSON(fiber.Map{
		"code":   200,
		"status": "OK",
		"data": model.UserResponse{
			Name:     userRequest.Name,
			Password: userRequest.Password,
			Token:    tokenString,
		},
	})
}
func SignUp(c *fiber.Ctx) error {
	db := db.Connect()
	userRequest := model.UserLogin{}
	c.BodyParser(&userRequest)

	user := model.User{
		Name:     userRequest.Name,
		Password: userRequest.Password,
	}

	if err := db.Debug().Model(&model.User{}).Create(&user).Error; err != nil {
		return c.Status(500).SendString("error creating user " + err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"code":   200,
		"status": "OK",
		"data": model.UserResponse{
			Name:     userRequest.Name,
			Password: userRequest.Password,
		},
	})
}

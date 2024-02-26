package main

import (
	"fiberjwt/handler"
	"fiberjwt/middleware"
	"fiberjwt/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New()
	app.Post("/login", handler.Login)
	app.Post("/sign-up", handler.SignUp)

	app.Use(csrf.New(csrf.Config{
		SingleUseToken: true,
		KeyLookup:      "header:CSRF_Token",
		CookieName:     "__Secure-csrf_",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		KeyGenerator:   utils.UUIDv4,
		Extractor:      csrf.CsrfFromHeader("CSRF_Token"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(err.Error())
		},
	}))

	app.Get("/user", middleware.Auth, func(c *fiber.Ctx) error {
		userRequest := model.UserLogin{}
		c.BodyParser(&userRequest)
		return c.Status(200).JSON(fiber.Map{
			"code":   200,
			"status": "OK",
			"data":   userRequest,
		})
	})

	app.Post("/post", middleware.Auth, handler.PostPost)
	app.Get("/post", handler.GetPost)
	app.Delete("/post/:id", middleware.Auth, handler.DeletePost)

	app.Listen(":8000")
}

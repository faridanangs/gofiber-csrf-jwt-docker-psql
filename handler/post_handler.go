package handler

import (
	"fiberjwt/db"
	"fiberjwt/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostPost(ctx *fiber.Ctx) error {
	db := db.Connect()
	request := model.Post{}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendString(err.Error())
	}
	request = model.Post{
		ID:       uuid.NewString(),
		Username: request.Username,
	}
	if err := db.Debug().Model(&model.Post{}).Create(&request).Error; err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(request)
}
func DeletePost(ctx *fiber.Ctx) error {
	db := db.Connect()
	id := ctx.Params("id")
	if err := db.Debug().Model(&model.Post{}).Where("id = ?", id).Delete(&model.Post{}).Error; err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.SendStatus(200)
}
func GetPost(ctx *fiber.Ctx) error {
	db := db.Connect()
	resp := []model.Post{}
	if err := db.Debug().Model(&model.Post{}).Find(&resp).Error; err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(resp)
}

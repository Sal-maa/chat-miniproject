package handler

import (
	"chat-miniproject/contract"
	"chat-miniproject/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateNewUser(ctx *fiber.Ctx) error {

	payload := new(contract.Users)

	err := ctx.BodyParser(payload)

	if err != nil {
		return err
	}

	user, errorInsert := service.CreateUser(*payload)
	if errorInsert != nil {
		return errorInsert
	}

	return ctx.JSON(fiber.Map{
		"inserted_record": user})
}

func GetAllUser(ctx *fiber.Ctx) error {
	users, err := service.FindAllUser()

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": 200,
		"data":   users})
}

func SendNewText(ctx *fiber.Ctx) error {

	payload := new(contract.Chat)

	err := ctx.BodyParser(payload)

	if err != nil {
		return err
	}

	text, errorInsert := service.SendText(*payload)
	if errorInsert != nil {
		return errorInsert
	}

	return ctx.JSON(fiber.Map{
		"inserted_record": text})
}

func ReadNewText(ctx *fiber.Ctx) error {
	text, err := service.ReadText()

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": 200,
		"data":   text})
}

func GetChatRoom(ctx *fiber.Ctx) error {
	// Id, errParseId := strconv.Atoi(ctx.Params("id"))
	// if errParseId != nil {
	// 	return errParseId
	// }

	var name string

	user, err := service.ChatRoom(name)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"data": user})
}

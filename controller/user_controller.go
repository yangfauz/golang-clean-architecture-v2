package controller

import (
	"api_project/exception"
	"api_project/model"
	"api_project/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{*userService}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Get("/api/users", controller.List)
	app.Get("/api/user/:id", controller.Detail)
	app.Post("/api/users", controller.Create)
	app.Put("/api/user/:id", controller.Update)
	app.Delete("/api/user/:id", controller.Delete)
}

func (controller *UserController) List(c *fiber.Ctx) error {
	responses := controller.UserService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserController) Detail(c *fiber.Ctx) error {
	var id string = c.Params("id")
	id_int, _ := strconv.Atoi(id)
	response := controller.UserService.Detail(id_int)
	if response.Id != 0 {
		return c.JSON(model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response,
		})
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "No Data",
		Data:   nil,
	})
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.UserService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Update(c *fiber.Ctx) error {
	var id string = c.Params("id")
	id_int, _ := strconv.Atoi(id)

	response := controller.UserService.Detail(id_int)
	if response.Id != 0 {
		var request model.CreateUserRequest
		err := c.BodyParser(&request)

		exception.PanicIfNeeded(err)

		response_update := controller.UserService.Update(id_int, request)

		return c.JSON(model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response_update,
		})
	}
	return c.Status(400).JSON(model.WebResponse{
		Code:   400,
		Status: "BAD_REQUEST",
		Data:   "Id: Not Found",
	})
}

func (controller *UserController) Delete(c *fiber.Ctx) error {
	var id string = c.Params("id")
	id_int, _ := strconv.Atoi(id)

	response := controller.UserService.Detail(id_int)
	if response.Id != 0 {
		response_delete := controller.UserService.Delete(id_int)
		return c.JSON(model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response_delete,
		})
	}
	return c.Status(400).JSON(model.WebResponse{
		Code:   400,
		Status: "BAD_REQUEST",
		Data:   "Id: Not Found",
	})
}

package user

import (
	"learn-portal/internal/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service domain.UserService
}

/*
	2xx - 299 -> menandakan proses berhasil
	4xx - 499 -> menandakan kesalahan dari sisi user
	5xx - n   -> menandakan kesalahan dari server
*/

func (i *Handler) Register(c *fiber.Ctx) error {
	body := domain.User{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := i.service.Register(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user register"})
}

// TODO : buat handler untuk read all user

func (i *Handler) GetAllUser(c *fiber.Ctx) error {
	users, err := i.service.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "all users succesfully retrieved",
		"users":   users,
	})

}

// TODO : buat handler untuk read one user by id menggunakan path parameter

func (i *Handler) GetUserProfile(c *fiber.Ctx) error {
	// Mengambil id dari http request
	// Konversi id tersebut ke integer
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	user, err := i.service.GetUserProfile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user succcesfully retrieved",
		"user":    user,
	})
}

func NewUserHandler(service domain.UserService) Handler {
	return Handler{service}
}

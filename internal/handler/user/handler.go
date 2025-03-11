package user

import (
	"learn-portal/internal/domain"

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

// TODO : buat handler untuk read one user by id menggunakan path parameter

func NewUserHandler(service domain.UserService) Handler {
	return Handler{service}
}

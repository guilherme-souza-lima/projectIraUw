package handler

import (
	"github.com/gofiber/fiber/v2"
	"projectIraUw/user_case/response"
)

type ServiceListINTER interface {
	ListAllService() ([]response.ListAll, error)
}

type ListHandler struct {
	ServiceListINTER ServiceListINTER
}

func NewListHandler(ServiceListINTER ServiceListINTER) ListHandler {
	return ListHandler{ServiceListINTER}
}

func (u ListHandler) ListAll(c *fiber.Ctx) error {
	result, err := u.ServiceListINTER.ListAllService()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

package hackerhandler

import (
	"github.com/gofiber/fiber/v2"
	"test_hackers/internal/core/ports"
)

type Handler struct {
	hackerService ports.HackerService
}

func NewHandler(hackerService ports.HackerService) *Handler {
	return &Handler{
		hackerService: hackerService,
	}
}

func (h Handler) GetAllHackers(c *fiber.Ctx) error {
	hackers, err := h.hackerService.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(hackers)
}

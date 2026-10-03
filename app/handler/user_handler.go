package handler

import (
	"errors"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/repository"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/service"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/helper"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.ErrorResponse(c, fiber.StatusNotFound, "user not found")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return helper.SuccessResponse(c, fiber.StatusOK, user)
}
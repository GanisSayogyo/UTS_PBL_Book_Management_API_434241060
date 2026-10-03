package handler

import (
	"errors"
	"strconv"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/model"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/repository"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/service"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnprocessableEntity, "validation failed")
	}

	user, err := h.authService.Register(c.Context(), req)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			return helper.ErrorResponse(c, fiber.StatusConflict, "username or email already exists")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	c.Location("/api/v1/users/" + strconv.Itoa(user.ID))

	return helper.SuccessResponse(c, fiber.StatusCreated, user)
}
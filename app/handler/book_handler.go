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

type BookHandler struct {
	bookService service.BookService
	validator   *validator.Validate
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
		validator:   validator.New(),
	}
}

func (h *BookHandler) GetAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid page")
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid limit")
	}

	search := c.Query("search", "")

	books, total, err := h.bookService.GetAll(
		c.Context(),
		page,
		limit,
		search,
	)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return helper.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"data": books,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (h *BookHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid book id")
	}

	book, err := h.bookService.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.ErrorResponse(c, fiber.StatusNotFound, "book not found")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return helper.SuccessResponse(c, fiber.StatusOK, book)
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	var req model.CreateBookRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnprocessableEntity, "validation failed")
	}

	book, err := h.bookService.Create(c.Context(), req)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			return helper.ErrorResponse(c, fiber.StatusConflict, "isbn already exists")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	c.Location("/api/v1/books/" + strconv.Itoa(book.ID))

	return helper.SuccessResponse(c, fiber.StatusCreated, book)
}

func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid book id")
	}

	var req model.UpdateBookRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnprocessableEntity, "validation failed")
	}

	book, err := h.bookService.Update(c.Context(), id, req)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.ErrorResponse(c, fiber.StatusNotFound, "book not found")
		}

		if errors.Is(err, repository.ErrDuplicate) {
			return helper.ErrorResponse(c, fiber.StatusConflict, "isbn already exists")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return helper.SuccessResponse(c, fiber.StatusOK, book)
}

func (h *BookHandler) Patch(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid book id")
	}

	var req model.PatchBookRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnprocessableEntity, "validation failed")
	}

	if req.Title == nil &&
		req.Author == nil &&
		req.ISBN == nil &&
		req.PublishedYear == nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "at least one field is required")
	}

	book, err := h.bookService.Patch(c.Context(), id, req)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.ErrorResponse(c, fiber.StatusNotFound, "book not found")
		}

		if errors.Is(err, repository.ErrDuplicate) {
			return helper.ErrorResponse(c, fiber.StatusConflict, "isbn already exists")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return helper.SuccessResponse(c, fiber.StatusOK, book)
}

func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "invalid book id")
	}

	if err := h.bookService.Delete(c.Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.ErrorResponse(c, fiber.StatusNotFound, "book not found")
		}

		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

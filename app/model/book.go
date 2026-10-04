package model

import "time"

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	ISBN          string    `json:"isbn"`
	PublishedYear int       `json:"published_year"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateBookRequest struct {
	Title         string `json:"title" validate:"required,min=2"`
	Author        string `json:"author" validate:"required,min=2"`
	ISBN          string `json:"isbn" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required"`
}

type UpdateBookRequest struct {
	Title         string `json:"title" validate:"required,min=2"`
	Author        string `json:"author" validate:"required,min=2"`
	ISBN          string `json:"isbn" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required"`
}

type PatchBookRequest struct {
	Title         *string `json:"title" validate:"omitempty,min=2"`
	Author        *string `json:"author" validate:"omitempty,min=2"`
	ISBN          *string `json:"isbn" validate:"omitempty"`
	PublishedYear *int    `json:"published_year" validate:"omitempty"`
}

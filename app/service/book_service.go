package service

import (
	"context"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/model"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/repository"
)

type BookService interface {
	Create(ctx context.Context, req model.CreateBookRequest) (*model.Book, error)
	GetAll(ctx context.Context, page, limit int, search string) ([]model.Book, int, error)
	GetByID(ctx context.Context, id int) (*model.Book, error)
	Update(ctx context.Context, id int, req model.UpdateBookRequest) (*model.Book, error)
	Patch(ctx context.Context, id int, req model.PatchBookRequest) (*model.Book, error)
	Delete(ctx context.Context, id int) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) Create(
	ctx context.Context,
	req model.CreateBookRequest,
) (*model.Book, error) {
	book := &model.Book{
		Title:         req.Title,
		Author:        req.Author,
		ISBN:          req.ISBN,
		PublishedYear: req.PublishedYear,
	}

	if err := s.bookRepo.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) GetAll(
	ctx context.Context,
	page, limit int,
	search string,
) ([]model.Book, int, error) {
	return s.bookRepo.FindAll(ctx, page, limit, search)
}

func (s *bookService) GetByID(
	ctx context.Context,
	id int,
) (*model.Book, error) {
	return s.bookRepo.FindByID(ctx, id)
}

func (s *bookService) Update(
	ctx context.Context,
	id int,
	req model.UpdateBookRequest,
) (*model.Book, error) {
	book := &model.Book{
		ID:            id,
		Title:         req.Title,
		Author:        req.Author,
		ISBN:          req.ISBN,
		PublishedYear: req.PublishedYear,
	}

	if err := s.bookRepo.Update(ctx, book); err != nil {
		return nil, err
	}

	return s.bookRepo.FindByID(ctx, id)
}

func (s *bookService) Delete(ctx context.Context, id int) error {
	return s.bookRepo.Delete(ctx, id)
}

func (s *bookService) Patch(ctx context.Context, id int, req model.PatchBookRequest) (*model.Book, error) {
	book, err := s.bookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		book.Title = *req.Title
	}

	if req.Author != nil {
		book.Author = *req.Author
	}

	if req.ISBN != nil {
		book.ISBN = *req.ISBN
	}

	if req.PublishedYear != nil {
		book.PublishedYear = *req.PublishedYear
	}

	if err := s.bookRepo.Update(ctx, book); err != nil {
		return nil, err
	}

	return s.bookRepo.FindByID(ctx, id)
}

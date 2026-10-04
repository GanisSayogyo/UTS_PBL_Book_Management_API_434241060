package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository interface {
	Create(ctx context.Context, book *model.Book) error
	FindAll(ctx context.Context, page, limit int, search string) ([]model.Book, int, error)
	FindByID(ctx context.Context, id int) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id int) error
}

type bookRepository struct {
	db *pgxpool.Pool
}

func NewBookRepository(db *pgxpool.Pool) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(ctx context.Context, book *model.Book) error {
	query := `
		INSERT INTO books (title, author, isbn, published_year)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		book.Title,
		book.Author,
		book.ISBN,
		book.PublishedYear,
	).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicate
		}

		return err
	}

	return nil
}

func (r *bookRepository) FindAll(
	ctx context.Context,
	page, limit int,
	search string,
) ([]model.Book, int, error) {
	offset := (page - 1) * limit

	countQuery := `
		SELECT COUNT(*)
		FROM books
		WHERE title ILIKE $1 OR author ILIKE $1 OR isbn ILIKE $1
	`

	var total int

	searchPattern := "%" + search + "%"

	if err := r.db.QueryRow(
		ctx,
		countQuery,
		searchPattern,
	).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, title, author, isbn, published_year, created_at, updated_at
		FROM books
		WHERE title ILIKE $1 OR author ILIKE $1 OR isbn ILIKE $1
		ORDER BY id
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(
		ctx,
		query,
		searchPattern,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	books := make([]model.Book, 0)

	for rows.Next() {
		var book model.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.PublishedYear,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (r *bookRepository) FindByID(ctx context.Context, id int) (*model.Book, error) {
	query := `
		SELECT id, title, author, isbn, published_year, created_at, updated_at
		FROM books
		WHERE id = $1
	`

	book := &model.Book{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.PublishedYear,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return book, nil
}

func (r *bookRepository) Update(ctx context.Context, book *model.Book) error {
	query := `
		UPDATE books
		SET title = $1,
			author = $2,
			isbn = $3,
			published_year = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		book.Title,
		book.Author,
		book.ISBN,
		book.PublishedYear,
		book.ID,
	).Scan(&book.UpdatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrDuplicate
		}

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

func (r *bookRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM books
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func buildBookSearchQuery(search string) string {
	if search == "" {
		return ""
	}

	return fmt.Sprintf("%%%s%%", search)
}
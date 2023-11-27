package book

import (
	"context"
	"fmt"

	"github.com/teamcubation/pod/internal/inventory/book/author"
	"github.com/teamcubation/pod/internal/platform/mysql"
)

type mysqlRepo struct {
	db *mysql.MySQLDb
}

func NewMySQLRepo(conn *mysql.MySQLDb) Repository {
	return &mysqlRepo{
		db: conn,
	}
}

func (r *mysqlRepo) Save(ctx context.Context, b *Book) error {
	bookDAO := mysql.BookDAO{
		Title:    b.Title,
		ISBN:     b.ISBN,
		Price:    b.Price,
		Stock:    b.Stock,
		AuthorID: b.Author.ID,
	}

	if err := r.db.SaveBook(ctx, &bookDAO); err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	b.ID = bookDAO.ID

	return nil
}

func (r *mysqlRepo) GetAll(ctx context.Context) ([]Book, error) {
	var books []Book

	booksDAO, err := r.db.GetBooks(ctx)
	if err != nil {
		return books, fmt.Errorf("error getting books: %w", err)
	}

	for _, b := range booksDAO {
		books = append(books, Book{
			ID: b.ID,
			Author: author.Author{
				ID:   b.AuthorID,
				Name: b.AuthorName,
				Age:  b.AuthorAge,
			},
			ISBN:  b.ISBN,
			Title: b.Title,
			Price: b.Price,
			Stock: b.Stock,
		})
	}

	return books, nil
}

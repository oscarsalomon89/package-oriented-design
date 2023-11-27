package book

import (
	"context"
	"fmt"

	"github.com/teamcubation/pod/internal/inventory/book/author"
)

type Repository interface {
	Save(ctx context.Context, book *Book) error
	GetAll(ctx context.Context) ([]Book, error)
}

type Book struct {
	ID     uint
	ISBN   string
	Title  string
	Price  float64
	Stock  int
	Author author.Author
}

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s no encontrado", e.Resource)
}

package author

import (
	"context"
	"errors"
)

type Repository interface {
	GetByID(ctx context.Context, id uint) (*Author, error)
	GetAll(ctx context.Context) ([]Author, error)
	Save(ctx context.Context, book *Author) error
}

type Author struct {
	ID   uint
	Name string
	Age  int
}

var ErrAuthorNotFound = errors.New("author not found")

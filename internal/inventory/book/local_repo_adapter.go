package book

import (
	"context"

	"github.com/teamcubation/pod/internal/inventory/book/author"
	"github.com/teamcubation/pod/internal/platform/memorydb"
)

var id = 0

type localRepo struct {
	db *memorydb.LocalDB
}

func NewLocalRepo(db *memorydb.LocalDB) Repository {
	return &localRepo{db: db}
}

func (l localRepo) Save(ctx context.Context, d *Book) error {
	id = id + 1

	book := memorydb.Book{
		AuthorID: d.Author.ID,
		Title:    d.Title,
		Price:    d.Price,
	}

	l.db.SaveItem(book)
	return nil
}

func (l localRepo) GetAll(ctx context.Context) ([]Book, error) {
	var bookSlice []Book

	items := l.db.GetAll()

	for _, b := range items {
		bookSlice = append(bookSlice, Book{
			ID:     b.ID,
			Author: author.Author{ID: b.AuthorID},
			Title:  b.Title,
			Price:  b.Price,
		})
	}

	return bookSlice, nil
}

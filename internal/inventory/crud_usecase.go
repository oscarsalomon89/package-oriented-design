package inventory

import (
	"context"
	"fmt"

	"github.com/teamcubation/pod/internal/inventory/book"
	"github.com/teamcubation/pod/internal/inventory/book/author"
)

// InputPort: UseCaseCRUD define la interfaz para el puerto de entrada del caso de uso.
type UseCaseCRUD interface {
	GetAllAuthors(ctx context.Context) ([]author.Author, error)
	GetAllBooks(ctx context.Context) ([]book.Book, error)
	SaveBook(ctx context.Context, book book.Book) (book.Book, error)
	SaveAuthor(ctx context.Context, name string, age int) (author.Author, error)
}

// Interactor: implementa la lógica de aplicación para los casos de uso de la biblioteca.
type useCaseCRUD struct {
	bookrepo   book.Repository
	authorrepo author.Repository
}

func NewUseCaseCRUD(bookrepo book.Repository, authorrepo author.Repository) *useCaseCRUD {
	return &useCaseCRUD{bookrepo: bookrepo, authorrepo: authorrepo}
}

func (c useCaseCRUD) SaveBook(ctx context.Context, newBook book.Book) (book.Book, error) {
	foundAuthor, err := c.authorrepo.GetByID(ctx, newBook.Author.ID)
	if err != nil {
		if err == author.ErrAuthorNotFound {
			return book.Book{}, &NotFoundError{Resource: "author"}
		}
		return book.Book{}, err
	}

	newBook.Author = *foundAuthor

	if err := c.bookrepo.Save(ctx, &newBook); err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (c useCaseCRUD) SaveAuthor(ctx context.Context, name string, age int) (author.Author, error) {
	author := author.Author{
		Name: name,
		Age:  age,
	}

	if err := c.authorrepo.Save(ctx, &author); err != nil {
		return author, err
	}

	return author, nil
}

func (c useCaseCRUD) GetAllBooks(ctx context.Context) ([]book.Book, error) {
	books, err := c.bookrepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// if books == nil {
	// 	return nil, fmt.Errorf("books not found")
	// }

	return books, nil
}

func (c useCaseCRUD) GetAllAuthors(ctx context.Context) ([]author.Author, error) {
	dev, err := c.authorrepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, fmt.Errorf("books not found")
	}

	return dev, nil
}

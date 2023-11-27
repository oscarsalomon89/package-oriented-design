package author

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *mysqlRepo) Save(ctx context.Context, a *Author) error {
	authorDAO := mysql.AuthorDAO{
		Name: a.Name,
		Age:  a.Age,
	}

	if err := r.db.SaveAuthor(ctx, &authorDAO); err != nil {
		return fmt.Errorf("error saving author: %w", err)
	}

	a.ID = authorDAO.ID

	return nil
}

func (r *mysqlRepo) GetByID(ctx context.Context, id uint) (*Author, error) {
	author, err := r.db.GetAuthorByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAuthorNotFound
		}
		return nil, fmt.Errorf("error getting author by ID: %w", err)
	}

	return &Author{
		ID:   author.ID,
		Name: author.Name,
		Age:  author.Age,
	}, nil
}

func (r *mysqlRepo) GetAll(ctx context.Context) ([]Author, error) {
	result, err := r.db.GetAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting authors: %w", err)
	}

	var authors []Author
	for _, b := range result {
		authors = append(authors, Author{
			ID:   b.ID,
			Name: b.Name,
			Age:  b.Age,
		})
	}

	return authors, nil
}

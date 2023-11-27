package app

import (
	"github.com/teamcubation/pod/internal/inventory/book"
	"github.com/teamcubation/pod/internal/inventory/book/author"
	"github.com/teamcubation/pod/internal/platform/environment"
	"github.com/teamcubation/pod/internal/platform/mysql"
)

type Dependencies struct {
	BookRepository   book.Repository
	AuthorRepository author.Repository
}

func BuildDependencies(env environment.Environment) (*Dependencies, error) {
	mysqlConn, err := mysql.GetConnectionDB()
	if err != nil {
		return nil, err
	}

	conn, err := mysql.NewMySQLDB(mysqlConn)
	if err != nil {
		return nil, err
	}

	bookRepo := book.NewMySQLRepo(conn)
	authorRepo := author.NewMySQLRepo(conn)

	return &Dependencies{
		BookRepository:   bookRepo,
		AuthorRepository: authorRepo,
	}, nil
}

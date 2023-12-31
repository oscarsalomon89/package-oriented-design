package app

import (
	"fmt"

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
	//localDb := memorydb.New()
	mysqlConn, err := mysql.GetConnectionDB()
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %w", err)
	}

	bookRepo := book.NewMySQLRepo(mysql.NewMySQLDB(mysqlConn))
	authorRepo := author.NewMySQLRepo(mysqlConn)
	//devRepo := book.NewLocalRepo(mysqlConn)

	return &Dependencies{
		BookRepository:   bookRepo,
		AuthorRepository: authorRepo,
	}, nil
}

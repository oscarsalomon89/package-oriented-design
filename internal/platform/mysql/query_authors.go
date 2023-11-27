package mysql

import (
	"context"
	"time"
)

type AuthorDAO struct {
	ID        uint
	Name      string
	Age       int
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (r *MySQLDb) SaveAuthor(ctx context.Context, a *AuthorDAO) error {
	createdAt := time.Now()
	result, err := r.conn.Exec(`INSERT INTO authors 
		(name, age, created_at, updated_at) 
		VALUES(?,?,?,?)`, a.Name, a.Age, createdAt, createdAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	a.ID = uint(id)

	return nil
}

func (r *MySQLDb) GetAuthorByID(ctx context.Context, id uint) (*AuthorDAO, error) {
	author := new(AuthorDAO)
	err := r.conn.Get(author, "SELECT * FROM authors WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (r *MySQLDb) GetAuthors(ctx context.Context) ([]AuthorDAO, error) {
	var authors []AuthorDAO

	err := r.conn.Select(&authors, `SELECT * FROM authors`)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

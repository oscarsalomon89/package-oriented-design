package author

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type authorDAO struct {
	ID        uint
	Name      string
	Age       int
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type mysqlRepo struct {
	conn *sqlx.DB
}

func NewMySQLRepo(conn *sqlx.DB) Repository {
	return &mysqlRepo{
		conn: conn,
	}
}

func (r *mysqlRepo) Save(ctx context.Context, a *Author) error {
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

func (r *mysqlRepo) GetByID(ctx context.Context, id uint) (*Author, error) {
	author := new(authorDAO)
	err := r.conn.Get(author, "SELECT * FROM authors WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAuthorNotFound
		}
		return nil, fmt.Errorf("error getting book: %w", err)
	}

	return &Author{
		ID:   author.ID,
		Name: author.Name,
		Age:  author.Age,
	}, nil
}

func (r *mysqlRepo) GetAll(ctx context.Context) ([]Author, error) {
	var authors []Author

	err := r.conn.Select(&authors, `SELECT * FROM authors`)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type BookDAO struct {
	ID         uint
	AuthorID   uint   `db:"author_id"`
	AuthorName string `db:"name"`
	AuthorAge  int    `db:"age"`
	ISBN       string `db:"isbn"`
	Title      string
	Price      float64
	Stock      int
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type MySQLDb struct {
	conn *sqlx.DB
}

func NewMySQLDB(conn *sqlx.DB) (*MySQLDb, error) {
	if conn == nil {
		return nil, errors.New("error initializing database")
	}
	return &MySQLDb{
		conn: conn,
	}, nil
}

func (r *MySQLDb) SaveBook(ctx context.Context, b *BookDAO) error {
	createdAt := time.Now()
	result, err := r.conn.Exec(`INSERT INTO books 
		(title, isbn, price, stock, author_id, created_at, updated_at) 
		VALUES(?,?,?,?,?,?,?)`, b.Title, b.ISBN, b.Price, b.Stock, b.AuthorID, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last inserted ID: %w", err)
	}

	b.CreatedAt = createdAt
	b.UpdatedAt = createdAt
	b.ID = uint(id)

	return nil
}

func (r *MySQLDb) GetBooks(ctx context.Context) ([]BookDAO, error) {
	var books []BookDAO

	err := r.conn.Select(&books, `SELECT * FROM books INNER JOIN authors ON books.author_id = authors.id`)
	if err != nil {
		return books, err
	}

	return books, nil
}

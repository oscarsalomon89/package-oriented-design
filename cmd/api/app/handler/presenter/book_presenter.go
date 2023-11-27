package presenter

import "github.com/teamcubation/pod/internal/inventory/book"

type jsonBook struct {
	ID     uint       `json:"id"`
	Author jsonAuthor `json:"author"`
	ISBN   string     `db:"isbn"`
	Title  string     `db:"title"`
	Price  float64    `db:"price"`
	Stock  int        `db:"stock"`
}

func Book(b book.Book) jsonBook {
	toReturn := jsonBook{
		ID:    b.ID,
		ISBN:  b.ISBN,
		Title: b.Title,
		Price: b.Price,
		Stock: b.Stock,
		Author: jsonAuthor{
			ID:   b.Author.ID,
			Name: b.Author.Name,
			Age:  b.Author.Age,
		},
	}

	return toReturn
}

func Books(books []book.Book) []jsonBook {
	jsonBooks := []jsonBook{}

	for _, b := range books {
		jsonBooks = append(jsonBooks, Book(b))
	}

	return jsonBooks
}

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

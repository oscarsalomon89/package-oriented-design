package presenter

import (
	"github.com/teamcubation/pod/internal/inventory/book/author"
)

type jsonAuthor struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Author(a author.Author) *jsonAuthor {
	toReturn := &jsonAuthor{
		ID:   a.ID,
		Name: a.Name,
		Age:  a.Age,
	}

	return toReturn
}

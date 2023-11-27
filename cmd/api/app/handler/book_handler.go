package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/pod/cmd/api/app/handler/presenter"
	"github.com/teamcubation/pod/internal/inventory"
	"github.com/teamcubation/pod/internal/inventory/book"
	"github.com/teamcubation/pod/internal/inventory/book/author"
)

type crudHandler struct {
	useCase inventory.UseCaseCRUD
}

func NewCRUDHandler(useCase inventory.UseCaseCRUD) *crudHandler {
	return &crudHandler{useCase: useCase}
}

type createBookRequest struct {
	Title    string  `json:"title" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	ISBN     string  `json:"isbn" binding:"required"`
	Stock    int     `json:"stock" binding:"required"`
	AuthorID uint    `json:"author_id" binding:"required"`
}

func (h crudHandler) HandleCreate(c *gin.Context) {
	var payload createBookRequest
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Los datos enviados son incorrectos: %s", err),
		})
		return
	}

	newBook := book.Book{
		Title:  payload.Title,
		ISBN:   payload.ISBN,
		Price:  payload.Price,
		Stock:  payload.Stock,
		Author: author.Author{ID: payload.AuthorID},
	}

	book, err := h.useCase.SaveBook(c, newBook)
	if err != nil {
		switch err := err.(type) {
		case *inventory.NotFoundError:
			c.JSON(http.StatusNotFound, presenter.ApiError{StatusCode: http.StatusNotFound, Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, presenter.ApiError{
			Message: "unexpected error",
		})
		return
	}

	c.JSON(http.StatusCreated, presenter.Book(book))
}

func (h crudHandler) HandleRead(c *gin.Context) {
	books, err := h.useCase.GetAllBooks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Books(books))
}

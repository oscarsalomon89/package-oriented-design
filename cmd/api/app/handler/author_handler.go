package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcubation/pod/cmd/api/app/handler/presenter"
	"github.com/teamcubation/pod/internal/inventory"
)

type authorHandler struct {
	useCase inventory.UseCaseCRUD
}

func NewAuthorHandler(useCase inventory.UseCaseCRUD) *authorHandler {
	return &authorHandler{useCase: useCase}
}

type createAuthorRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func (h authorHandler) HandleCreate(c *gin.Context) {
	var payload createAuthorRequest
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{})
		return
	}

	result, err := h.useCase.SaveAuthor(c, payload.Name, payload.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusCreated, presenter.Author(result))
}

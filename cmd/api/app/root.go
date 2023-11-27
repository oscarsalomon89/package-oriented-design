package app

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcubation/pod/cmd/api/app/handler"
	"github.com/teamcubation/pod/internal/inventory"
)

func Build(dep *Dependencies) *gin.Engine {
	router := gin.Default()

	// use cases
	inventoryUseCase := inventory.NewUseCaseCRUD(dep.BookRepository, dep.AuthorRepository)

	// controller adapters
	bookHandler := handler.NewCRUDHandler(inventoryUseCase)
	authorHandler := handler.NewAuthorHandler(inventoryUseCase)

	basePath := "/api/v1/inventory"
	r := router.Group(basePath)

	r.GET("/books", bookHandler.HandleRead)
	r.POST("/books", bookHandler.HandleCreate)
	r.POST("/authors", authorHandler.HandleCreate)

	return router
}

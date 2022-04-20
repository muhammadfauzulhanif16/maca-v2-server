package routes

import (
	"maca/config"
	"maca/handlers"
	"maca/repositories"
	"maca/services"

	"github.com/gin-gonic/gin"
)

var (
	db    = config.ConnectDb()
	bookR = repositories.NewBookR(db)
	bookS = services.NewBookS(bookR)
	bookH = handlers.NewBookH(bookS)
)

func Book(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	book := v1.Group("/book")
	book.POST("", bookH.Create)
	book.PUT("/:id", bookH.UpdateIsCompleted)
	book.DELETE("/:id", bookH.Delete)

	books := v1.Group("/books")
	books.GET("", bookH.ReadAll)
	books.DELETE("", bookH.DeleteAll)
	books.GET("/search", bookH.ReadSearch)
}

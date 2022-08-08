package routes

import (
	"book-app/controller"
	"github.com/gin-gonic/gin"
)

func BookRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("books/create", controller.CreateBook())
	incomingRoutes.GET("books/:book_id", controller.GetBook())
	incomingRoutes.PATCH("books/:book_id", controller.UpdateBook())
	incomingRoutes.DELETE("books/:book_id", controller.DeleteBook())
	incomingRoutes.GET("books", controller.GetAllBooks())
}

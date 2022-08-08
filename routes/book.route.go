package routes

import (
	"book-app/controller"
	"github.com/gin-gonic/gin"
)

func BookRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("books/create", controller.CreateBook())
	// incomingRoutes.GET("books/:book_id", controllers.GetBook())
	// incomingRoutes.PATCH("books/:book_id", controllers.UpdateBook())
	// incomingRoutes.DELETE("books/:book_id", controllers.DeleteBook())
	// incomingRoutes.GET("books", controllers.GetAllBooks())
}

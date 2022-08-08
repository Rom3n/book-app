package controller

import (
	"book-app/database"
	"book-app/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var bookCollection = database.OpenConnection(database.Client, "book")
var validate = validator.New()

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var book model.Book

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(book)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		book.Created, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		book.Updated, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		book.ID = primitive.NewObjectID()

		result, insertError := bookCollection.InsertOne(ctx, book)
		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Book item wasn't created."})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

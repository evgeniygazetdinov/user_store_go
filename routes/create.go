package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	getcollection "work_in_que/Collection"
	database "work_in_que/database"
	model "work_in_que/model"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var DB = database.ConnectDB()
	fmt.Println("before")
	var postCollection = getcollection.GetCollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("herer")
	post := new(model.Posts)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)

		return
	}

	postPayload := model.Posts{
		Title:   post.Title,
		Article: post.Article,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}

package main

import (
	"net/http"

	"github.com/bhavyaunacademy/goCRUD-MySQL/database"
	handler "github.com/bhavyaunacademy/goCRUD-MySQL/handlers"
	"github.com/bhavyaunacademy/goCRUD-MySQL/models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//Intialising the db
	db := database.InitDb()
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	//Intialising the handler
	h := &handler.Handler{DB: db}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// routes for getting users
	e.GET("/users", h.GetAllUsers)
	e.GET("/users/:uid", h.GetUserByUID)
	e.POST("/users", h.CreateUser)

	// routes for getting posts
	e.GET("/posts", h.GetAllPosts)
	e.GET("/posts/:uid", h.GetPostByUID)
	e.GET("/posts/user/:uid", h.GetPostByUserUID)

	// CRUD routes for Post
	e.POST("/posts/create-post", h.CreatePost)
	e.PUT("/posts/:postUID/update-post", h.UpdatePost)
	e.DELETE("/posts/:postUID/delete-post", h.DeletePost)

	// Like/DisLike on Post
	e.PUT("/posts/:postUID/like", h.LikePost)
	e.PUT("/posts/:postUID/dislike", h.DislikePost)

	// routes for getting posts
	e.GET("/posts/:postUID/comments", h.GetCommentsOnAPost)
	// CRUD routes for Comments
	e.POST("/comments/create-comment", h.CreateComment)
	e.PUT("/comments/:commentUID/update-comment", h.UpdateComment)
	e.DELETE("/comments/:commentUID/delete-comment", h.DeleteComment)

	e.Logger.Fatal(e.Start(":1323"))
}

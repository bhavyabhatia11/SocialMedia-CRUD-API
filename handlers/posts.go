package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bhavyaunacademy/goCRUD-MySQL/models"
	"github.com/bhavyaunacademy/goCRUD-MySQL/modules"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) GetAllPosts(c echo.Context) error {
	var postList []models.Post

	err := modules.GetPosts(h.DB, &postList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var posts models.Posts
	posts.Posts = postList

	return c.JSON(http.StatusCreated, posts)
}

func (h *Handler) GetPostByUID(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var post models.Post
	err = modules.GetPost(h.DB, &post, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)

}

func (h *Handler) GetPostByUserUID(c echo.Context) error {
	user_uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var postList []models.Post
	err = modules.GetPostsByUser(h.DB, &postList, user_uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var posts models.Posts
	posts.Posts = postList

	return c.JSON(http.StatusOK, posts)
}

func (h *Handler) CreatePost(c echo.Context) (err error) {

	post := new(models.Post)
	if err = c.Bind(post); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = modules.CreatePost(h.DB, post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, post)
}

func (h *Handler) UpdatePost(c echo.Context) (err error) {

	post := new(models.Post)
	if err = c.Bind(post); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = modules.UpdatePost(h.DB, post, post_uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, post)

}

func (h *Handler) DeletePost(c echo.Context) (err error) {

	post := new(models.Post)
	if err = c.Bind(post); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = modules.DeletePost(h.DB, post, post_uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, post)
}

func (h *Handler) LikePost(c echo.Context) (err error) {
	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = modules.LikePost(h.DB, post_uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "Post Liked")
}

func (h *Handler) DislikePost(c echo.Context) (err error) {
	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = modules.DislikePost(h.DB, post_uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "Post Disliked")
}

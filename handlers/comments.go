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

func (h *Handler) GetCommentsOnAPost(c echo.Context) (err error) {
	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var commentList []models.Comment
	err = modules.GetCommentsOfPost(h.DB, &commentList, post_uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var comments models.Comments
	comments.Comments = commentList

	return c.JSON(http.StatusOK, comments)
}

func (h *Handler) CreateComment(c echo.Context) (err error) {

	comment := new(models.Comment)
	if err = c.Bind(comment); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = modules.CreateComment(h.DB, comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, comment)
}

func (h *Handler) UpdateComment(c echo.Context) (err error) {

	comment_uid, err := strconv.Atoi(c.Param("commentUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	comment := new(models.Comment)
	if err = c.Bind(comment); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = modules.UpdateComment(h.DB, comment, comment_uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, comment)
}

func (h *Handler) DeleteComment(c echo.Context) (err error) {

	comment := new(models.Comment)
	if err = c.Bind(comment); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	comment_uid, err := strconv.Atoi(c.Param("commentUID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = modules.DeleteComment(h.DB, comment, comment_uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, comment)
}

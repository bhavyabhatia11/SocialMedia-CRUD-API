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

func (h *Handler) GetAllUsers(c echo.Context) (err error) {
	var userList []models.User

	err = modules.GetUsers(h.DB, &userList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var users models.Users
	users.Users = userList

	return c.JSON(http.StatusCreated, users)
}

func (h *Handler) GetUserByUID(c echo.Context) (err error) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var user models.User
	err = modules.GetUser(h.DB, &user, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	return c.JSON(http.StatusOK, user)

}

func (h *Handler) CreateUser(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = modules.CreateUser(h.DB, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

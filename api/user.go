package api

import (
	"github.com/iamyusuf/gws/types/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	s.Db.Create(&user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (s *Server) FindUserById(c echo.Context) error {
	id := c.Param("id")
	user := model.User{}
	s.Db.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

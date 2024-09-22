package api

import (
	"github.com/iamyusuf/gws/types/model"
	"github.com/iamyusuf/gws/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) CreateUser(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user.Password = hashedPassword

	if result := s.Db.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not save user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func (s *Server) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User
	s.Db.First(&user, id)

	err := c.Bind(&user)

	if err != nil {
		return err
	}

	s.Db.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func (s *Server) FindUserById(c echo.Context) error {
	id := c.Param("id")
	user := model.User{}
	s.Db.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

func (s *Server) DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	s.Db.Delete(&model.User{}, id)
	return c.JSON(http.StatusOK, nil)
}

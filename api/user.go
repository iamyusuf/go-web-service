package api

import (
	"github.com/iamyusuf/gws/types/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	s.Db.Create(&user)

	if err != nil {
		return err
	}

	return nil
}

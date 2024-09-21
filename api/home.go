package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Echo Sever. Version 1.0.0")
}

package server

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, generated.HelloResponse{
		Message: "Hello, 🌏! ദ്ദി(⎚_⎚)",
	})
}

func (s *Server) PostEstate(ctx echo.Context) error {
	return s.estateHandler.PostEstate(ctx)
}

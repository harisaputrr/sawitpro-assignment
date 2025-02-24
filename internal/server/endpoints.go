package server

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
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

func (s *Server) PostEstateIdTree(ctx echo.Context, uuid uuid.UUID) error {
	return s.treeHandler.PostTree(ctx, uuid)
}

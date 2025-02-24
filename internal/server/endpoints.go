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

func (s *Server) GetEstateStats(ctx echo.Context, uuid uuid.UUID) error {
	return s.estateHandler.GetEstateStats(ctx, uuid)
}

func (s *Server) CreateEstate(ctx echo.Context) error {
	return s.estateHandler.CreateEstate(ctx)
}

func (s *Server) CreateTree(ctx echo.Context, uuid uuid.UUID) error {
	return s.treeHandler.CreateTree(ctx, uuid)
}

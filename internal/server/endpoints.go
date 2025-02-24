package server

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateStats(ctx echo.Context, uuid uuid.UUID) error {
	return s.estateHandler.GetEstateStats(ctx, uuid)
}

func (s *Server) GetEstateDronePlan(ctx echo.Context, uuid uuid.UUID, params generated.GetEstateDronePlanParams) error {
	return s.estateHandler.GetEstateDronePlan(ctx, uuid, params)
}

func (s *Server) CreateEstate(ctx echo.Context) error {
	return s.estateHandler.CreateEstate(ctx)
}

func (s *Server) CreateTree(ctx echo.Context, uuid uuid.UUID) error {
	return s.treeHandler.CreateTree(ctx, uuid)
}

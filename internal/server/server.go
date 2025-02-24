package server

import (
	estateHandler "github.com/SawitProRecruitment/UserService/internal/estate/handler"
	treeHandler "github.com/SawitProRecruitment/UserService/internal/tree/handler"
)

type Server struct {
	estateHandler *estateHandler.EstateHandler
	treeHandler   *treeHandler.TreeHandler
}

func NewServer(
	estateHandler *estateHandler.EstateHandler,
	treeHandler *treeHandler.TreeHandler,
) *Server {
	return &Server{
		estateHandler: estateHandler,
		treeHandler:   treeHandler,
	}
}

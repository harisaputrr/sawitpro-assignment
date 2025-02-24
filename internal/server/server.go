package server

import (
	estateHandler "github.com/SawitProRecruitment/UserService/internal/estate/handler"
)

type Server struct {
	estateHandler *estateHandler.EstateHandler
}

func NewServer(
	estateHandler *estateHandler.EstateHandler,
) *Server {
	return &Server{
		estateHandler: estateHandler,
	}
}

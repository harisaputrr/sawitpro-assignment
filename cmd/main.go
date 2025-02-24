package main

import (
	"database/sql"

	"github.com/SawitProRecruitment/UserService/generated"
	estateHandler "github.com/SawitProRecruitment/UserService/internal/estate/handler"
	estateRepository "github.com/SawitProRecruitment/UserService/internal/estate/repository"
	estateUsecase "github.com/SawitProRecruitment/UserService/internal/estate/usecase"
	"github.com/SawitProRecruitment/UserService/internal/server"
	treeHandler "github.com/SawitProRecruitment/UserService/internal/tree/handler"
	treeRepository "github.com/SawitProRecruitment/UserService/internal/tree/repository"
	treeUsecase "github.com/SawitProRecruitment/UserService/internal/tree/usecase"
	"github.com/SawitProRecruitment/UserService/pkg/config"
	"github.com/SawitProRecruitment/UserService/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.LoadConfig()
	db := database.NewPostgre(conf)

	server := newServer(db)

	e := echo.New()

	e.Use(middleware.Logger())

	generated.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":" + conf.AppPort))
}

func newServer(db *sql.DB) *server.Server {
	estateRepository := estateRepository.NewRepository(db)
	treeRepository := treeRepository.NewRepository(db)

	estateUsecase := estateUsecase.NewUsecase(estateRepository, treeRepository)
	treeUsecase := treeUsecase.NewUsecase(treeRepository, estateRepository)

	estateHandler := estateHandler.NewHandler(estateUsecase)
	treeHandler := treeHandler.NewHandler(treeUsecase)

	return server.NewServer(estateHandler, treeHandler)
}

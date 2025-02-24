package database

import (
	"fmt"

	"database/sql"

	"github.com/SawitProRecruitment/UserService/pkg/config"
	_ "github.com/lib/pq"
)

func NewPostgre(conf *config.Config) *sql.DB {
	db, err := sql.Open("postgres", conf.PostgreDSN)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}

	return db
}

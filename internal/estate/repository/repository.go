package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/SawitProRecruitment/UserService/generated"
)

type Repository struct {
	db    *sql.DB
	table string
}

func NewRepository(db *sql.DB) EstateRepository {
	return &Repository{
		db:    db,
		table: "estates",
	}
}

func (r *Repository) CreateEstate(ctx context.Context, input generated.CreateEstateRequest) (output *generated.CreateEstateResponse, err error) {

	if input.Width < 1 || input.Width > 50000 {
		return output, errors.New("length must be between 1 and 50000")
	}
	if input.Length < 1 || input.Length > 50000 {
		return output, errors.New("length must be between 1 and 50000")
	}

	output = new(generated.CreateEstateResponse)

	query := fmt.Sprintf("INSERT INTO %s (width, length) VALUES ($1, $2) RETURNING id", r.table)
	err = r.db.QueryRowContext(ctx, query, input.Width, input.Length).Scan(&output.Id)
	if err != nil {
		return output, fmt.Errorf("failed to create estate: %w", err)
	}

	return output, nil
}

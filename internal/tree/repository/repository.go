package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/tree/model"
	"github.com/google/uuid"
)

type Repository struct {
	db    *sql.DB
	table string
}

func NewRepository(db *sql.DB) TreeRepository {
	return &Repository{
		db:    db,
		table: "Trees",
	}
}

func (r *Repository) CreateTree(ctx context.Context, estateID uuid.UUID, input generated.CreateTreeRequest) (output *generated.CreateTreeResponse, err error) {
	output = new(generated.CreateTreeResponse)

	query := `
		INSERT INTO trees (estate_id, x, y, height, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err = r.db.QueryRow(query, estateID, input.X, input.Y, input.Height, time.Now()).Scan(&output.Id)
	if err != nil {
		return output, fmt.Errorf("failed to create tree: %w", err)
	}

	return output, nil
}

func (r *Repository) FindOneTreebyCoordinate(ctx context.Context, estateID uuid.UUID, x, y int) (t *model.Tree, err error) {
	t = new(model.Tree)

	query := `
		SELECT *
		FROM trees
		WHERE estate_id = $1 AND x = $2 AND y = $3
	`

	err = r.db.QueryRow(query, estateID, x, y).Scan(
		&t.ID, &t.EstateID, &t.X, &t.Y, &t.Height, &t.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return t, nil
}

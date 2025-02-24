package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/tree/model"
	"github.com/google/uuid"
)

type TreeRepository interface {
	CreateTree(ctx context.Context, estateID uuid.UUID, input generated.CreateTreeRequest) (output *generated.CreateTreeResponse, err error)
	FindOneTreebyCoordinate(ctx context.Context, estateID uuid.UUID, x, y int) (t *model.Tree, err error)
}

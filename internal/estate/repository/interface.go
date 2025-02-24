package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/estate/model"
	"github.com/google/uuid"
)

type EstateRepository interface {
	GetEstateByID(ctx context.Context, estateID uuid.UUID) (output *model.Estate, err error)
	GetEstateStats(ctx context.Context, estateID uuid.UUID) (output *generated.EstateStatsResponse, err error)
	CreateEstate(ctx context.Context, input generated.CreateEstateRequest) (output *generated.CreateEstateResponse, err error)
}

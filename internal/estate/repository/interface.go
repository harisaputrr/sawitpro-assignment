package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
)

type EstateRepository interface {
	GetEstateStats(ctx context.Context, estateID uuid.UUID) (output *generated.EstateStatsResponse, err error)
	CreateEstate(ctx context.Context, input generated.CreateEstateRequest) (output *generated.CreateEstateResponse, err error)
}

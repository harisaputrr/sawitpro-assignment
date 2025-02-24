package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
)

type EstateUsecase interface {
	GetEstateStats(ctx context.Context, estateID uuid.UUID) (result *generated.EstateStatsResponse, err error)
	CreateEstate(ctx context.Context, payload generated.CreateEstateRequest) (result *generated.CreateEstateResponse, err error)
}

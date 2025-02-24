package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/estate/repository"
	"github.com/google/uuid"
)

type Usecase struct {
	repository repository.EstateRepository
}

func NewUsecase(repository repository.EstateRepository) EstateUsecase {
	return &Usecase{repository: repository}
}

func (u *Usecase) GetEstateStats(ctx context.Context, estateID uuid.UUID) (result *generated.EstateStatsResponse, err error) {
	result, err = u.repository.GetEstateStats(ctx, estateID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *Usecase) CreateEstate(ctx context.Context, payload generated.CreateEstateRequest) (result *generated.CreateEstateResponse, err error) {
	result, err = u.repository.CreateEstate(ctx, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

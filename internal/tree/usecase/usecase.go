package usecase

import (
	"context"
	"errors"

	"github.com/SawitProRecruitment/UserService/generated"
	estateRepository "github.com/SawitProRecruitment/UserService/internal/estate/repository"
	"github.com/SawitProRecruitment/UserService/internal/tree/repository"
	"github.com/google/uuid"
)

type Usecase struct {
	repository       repository.TreeRepository
	estateRepository estateRepository.EstateRepository
}

func NewUsecase(repository repository.TreeRepository, estateRepository estateRepository.EstateRepository) TreeUsecase {
	return &Usecase{
		repository:       repository,
		estateRepository: estateRepository,
	}
}

func (u *Usecase) CreateTree(ctx context.Context, estateID uuid.UUID, payload generated.CreateTreeRequest) (result *generated.CreateTreeResponse, err error) {
	estate, err := u.estateRepository.GetEstateByID(ctx, estateID)
	if err != nil {
		return nil, err
	}

	if payload.X < 1 || payload.X > estate.Length || payload.Y < 1 || payload.Y > estate.Width {
		return nil, errors.New("coordinates are out of estate bounds")
	}

	existingTree, err := u.repository.FindOneTreebyCoordinate(ctx, estateID, payload.X, payload.Y)
	if err != nil {
		return nil, err
	}
	if existingTree != nil {
		return nil, errors.New("a tree already exists at the given coordinates")
	}

	result, err = u.repository.CreateTree(ctx, estateID, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

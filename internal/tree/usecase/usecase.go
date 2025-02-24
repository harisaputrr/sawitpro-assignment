package usecase

import (
	"context"
	"errors"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/tree/repository"
	"github.com/google/uuid"
)

type Usecase struct {
	repository repository.TreeRepository
}

func NewUsecase(repository repository.TreeRepository) TreeUsecase {
	return &Usecase{repository: repository}
}

func (u *Usecase) CreateTree(ctx context.Context, estateID uuid.UUID, payload generated.CreateTreeRequest) (result *generated.CreateTreeResponse, err error) {
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

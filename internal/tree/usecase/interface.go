package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
)

type TreeUsecase interface {
	CreateTree(ctx context.Context, estateID uuid.UUID, payload generated.CreateTreeRequest) (result *generated.CreateTreeResponse, err error)
}

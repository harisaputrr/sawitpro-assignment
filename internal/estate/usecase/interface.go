package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
)

type EstateUsecase interface {
	CreateEstate(ctx context.Context, payload generated.CreateEstateRequest) (result *generated.CreateEstateResponse, err error)
}

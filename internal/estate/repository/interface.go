package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
)

type EstateRepository interface {
	CreateEstate(ctx context.Context, input generated.CreateEstateRequest) (output *generated.CreateEstateResponse, err error)
}

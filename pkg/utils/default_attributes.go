package utils

import (
	"time"

	"github.com/google/uuid"
)

type DefaultAttributes struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type DefaultPaginationAttributes struct {
	Page              int  `json:"page" query:"page"`
	Limit             int  `json:"limit" query:"limit" validate:"min=1"`
	DisablePagination bool `json:"disable_pagination" query:"disable_pagination"`
}

func (d *DefaultPaginationAttributes) CalculateOffset() {
	if d.DisablePagination {
		d.Page = 0
	}
	d.Page = d.Limit * (d.Page - 1)
}

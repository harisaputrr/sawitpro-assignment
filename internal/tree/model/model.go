package model

import (
	"github.com/SawitProRecruitment/UserService/pkg/utils"
	"github.com/google/uuid"
)

type (
	Tree struct {
		utils.DefaultAttributes
		EstateID uuid.UUID `json:"estate_id"`
		X        int       `json:"x"`
		Y        int       `json:"y"`
		Height   int       `json:"height"`
	}
)

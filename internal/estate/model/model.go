package model

import (
	"github.com/SawitProRecruitment/UserService/pkg/utils"
)

type (
	Estate struct {
		utils.DefaultAttributes
		Width  int `json:"width"`
		Length int `json:"length"`
	}
)

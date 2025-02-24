package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/estate/usecase"
	"github.com/labstack/echo/v4"
)

type EstateHandler struct {
	usecase usecase.EstateUsecase
}

func NewHandler(estateUsecase usecase.EstateUsecase) *EstateHandler {
	return &EstateHandler{usecase: estateUsecase}
}

func (h *EstateHandler) PostEstate(ctx echo.Context) error {
	var payload generated.CreateEstateRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.usecase.CreateEstate(ctx.Request().Context(), payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, result)
}

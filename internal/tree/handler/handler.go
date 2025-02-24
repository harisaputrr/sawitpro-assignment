package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/internal/tree/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TreeHandler struct {
	usecase usecase.TreeUsecase
}

func NewHandler(treeUsecase usecase.TreeUsecase) *TreeHandler {
	return &TreeHandler{usecase: treeUsecase}
}

func (h *TreeHandler) CreateTree(ctx echo.Context, estateID uuid.UUID) error {
	var payload generated.CreateTreeRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.usecase.CreateTree(ctx.Request().Context(), estateID, payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, result)
}

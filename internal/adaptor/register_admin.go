package adaptor

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/internal/usecase"
	"project-POS-APP-golang-be-team/pkg/response"

	"github.com/gin-gonic/gin"
)

type RegisterAdminHandler struct {
	UC usecase.RegisterAdminUsecase
}

func NewRegisterAdminHandler(uc usecase.RegisterAdminUsecase) *RegisterAdminHandler {
	return &RegisterAdminHandler{UC: uc}
}

func (h *RegisterAdminHandler) RegisterAdmin(c *gin.Context) {
	var req dto.RegisterAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseBadRequest(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.UC.RegisterAdmin(c.Request.Context(), req); err != nil {
		response.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ResponseSuccess(c, http.StatusCreated, "Admin registered successfully", nil)
}

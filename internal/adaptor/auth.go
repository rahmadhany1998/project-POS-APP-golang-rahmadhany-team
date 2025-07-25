package adaptor

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/internal/usecase"
	"project-POS-APP-golang-be-team/pkg/response"
	"project-POS-APP-golang-be-team/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerAuth struct {
	AuthService usecase.AuthService
	Logger      *zap.Logger
}

func NewHandlerAuth(auth usecase.AuthService, logger *zap.Logger) HandlerAuth {
	return HandlerAuth{
		AuthService: auth,
		Logger:      logger,
	}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := utils.ValidateDataGin(err)
		response.ResponseBadRequest2(ctx, http.StatusBadRequest, detail)
		return
	}

	resp, err := h.AuthService.Login(ctx.Request.Context(), req)
	if err != nil {
		response.ResponseBadRequest(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	response.ResponseSuccess(ctx, http.StatusOK, "login successfull", resp)
}

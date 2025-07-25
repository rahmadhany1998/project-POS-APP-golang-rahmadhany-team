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

func (h *HandlerAuth) ForgotPassword(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := utils.ValidateDataGin(err)
		response.ResponseBadRequest2(ctx, http.StatusBadRequest, detail)
		return
	}
	if err := h.AuthService.ForgotPassword(ctx.Request.Context(), req); err != nil {
		response.ResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "otp sent!", nil)
}

func (h *HandlerAuth) VerifyOtp(ctx *gin.Context) {
	var req dto.VerifyOtpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := utils.ValidateDataGin(err)
		response.ResponseBadRequest2(ctx, http.StatusBadRequest, detail)
		return
	}
	if err := h.AuthService.VerifyOtp(ctx.Request.Context(), req); err != nil {
		response.ResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "otp valid", nil)
}

func (h *HandlerAuth) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := utils.ValidateDataGin(err)
		response.ResponseBadRequest2(ctx, http.StatusBadRequest, detail)
		return
	}
	if err := h.AuthService.ResetPassword(ctx.Request.Context(), req); err != nil {
		response.ResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "reset password success", nil)
}

func (h *HandlerAuth) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		response.ResponseBadRequest(ctx, http.StatusUnauthorized, "empty token")
		return
	}

	err := h.AuthService.Logout(ctx.Request.Context(), token)
	if err != nil {
		response.ResponseBadRequest(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.ResponseSuccess(ctx, http.StatusOK, "logout successfull", nil)
}

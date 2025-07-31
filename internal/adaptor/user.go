package adaptor

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/internal/usecase"
	"project-POS-APP-golang-be-team/pkg/response"

	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerUser struct {
	User   usecase.UserService
	Logger *zap.Logger
}

func NewHandlerUser(user usecase.UserService, logger *zap.Logger) HandlerUser {
	return HandlerUser{
		User:   user,
		Logger: logger,
	}
}

func (h *HandlerUser) TestHandler(ctx *gin.Context) {
	response.ResponseSuccess(ctx, http.StatusOK, "Ini adalah test handler", nil)
}

func (h *HandlerUser) GetProfile(ctx *gin.Context) {
	userIDStr := ctx.GetString("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	profile, err := h.User.GetProfile(userID)
	if err != nil {
		response.ResponseError(ctx, http.StatusInternalServerError, "Gagal mengambil profil")
		return
	}

	response.ResponseSuccess(ctx, http.StatusOK, "Berhasil mengambil profil", profile)
}

func (h *HandlerUser) UpdateProfile(ctx *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, "Data tidak valid")
		return
	}

	userIDStr := ctx.GetString("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	err = h.User.UpdateProfile(userID, req)
	if err != nil {
		response.ResponseError(ctx, http.StatusInternalServerError, "Gagal update profil")
		return
	}

	response.ResponseSuccess(ctx, http.StatusOK, "Profil berhasil diperbarui", nil)
}

func (h *HandlerUser) GetAdminList(ctx *gin.Context) {
	admins, err := h.User.GetAdminList()
	if err != nil {
		response.ResponseError(ctx, http.StatusInternalServerError, "Gagal mengambil daftar admin")
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "Berhasil mengambil daftar admin", admins)
}

func (h *HandlerUser) UpdateAdminAccess(ctx *gin.Context) {
	var req dto.UpdateAdminAccessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, "Data tidak valid")
		return
	}

	superAdminIDStr := ctx.GetString("user_id")
	superAdminID, err := strconv.Atoi(superAdminIDStr)
	if err != nil {
		response.ResponseError(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	err = h.User.UpdateAdminAccess(superAdminID, req)
	if err != nil {
		response.ResponseError(ctx, http.StatusForbidden, err.Error())
		return
	}

	response.ResponseSuccess(ctx, http.StatusOK, "Akses admin berhasil diperbarui", nil)
}

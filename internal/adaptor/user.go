package adaptor

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/usecase"
	"project-POS-APP-golang-be-team/pkg/response"

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

// Test handler (dummy)
func (h *HandlerUser) TestHandler(ctx *gin.Context) {
	response.ResponseSuccess(ctx, http.StatusOK, "Ini adalah test handler", nil)
}

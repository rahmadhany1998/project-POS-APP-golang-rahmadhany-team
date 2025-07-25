package middleware

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	Repo   repository.Repository
	Logger *zap.Logger
}

func NewAuthMiddleware(repo repository.Repository, logger *zap.Logger) AuthMiddleware {
	return AuthMiddleware{
		Repo:   repo,
		Logger: logger,
	}
}

func (m *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			response.ResponseBadRequest(ctx, http.StatusUnauthorized, "login first")
			ctx.Abort()
			return
		}

		// Ambil user berdasarkan token login
		loginToken, err := m.Repo.AuthRepo.FindUserByToken(ctx.Request.Context(), token)
		if err != nil {
			m.Logger.Error("token invalid", zap.String("token", token), zap.String("error", err.Error()))
			response.ResponseBadRequest(ctx, http.StatusUnauthorized, "invalid token")
			ctx.Abort()
			return
		}

		// Set ke context
		ctx.Set("userID", loginToken.User.ID)
		ctx.Set("userRole", loginToken.User.Role)

		ctx.Next()
	}
}

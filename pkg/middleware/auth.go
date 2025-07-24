package middleware

import (
	"fmt"
	"net/http"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	Repo   *gorm.DB
	Logger *zap.Logger
}

func NewAuthMiddleware(repo *gorm.DB, logger *zap.Logger) AuthMiddleware {
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

		var user entity.User
		err := m.Repo.Where("token = ?", token).First(&user).Error
		if err != nil {
			response.ResponseBadRequest(ctx, http.StatusUnauthorized, "invalid token")
			ctx.Abort()
			return
		}

		ctx.Set("userID", user.ID)
		ctx.Set("userRole", user.Role)
		fmt.Println("✅ Middleware: user found, ID =", user.ID, "Role =", user.Role)

		ctx.Next()
	}
}

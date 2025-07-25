package wire

import (
	"project-POS-APP-golang-be-team/internal/adaptor"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/usecase"
	"project-POS-APP-golang-be-team/pkg/middleware"
	"project-POS-APP-golang-be-team/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Wiring(repo repository.Repository, mLogger middleware.LoggerMiddleware, middlwareAuth middleware.AuthMiddleware, logger *zap.Logger, config utils.Configuration) *gin.Engine {
	router := gin.New()
	router.Use(mLogger.LoggingMiddleware())
	api := router.Group("/api/v1")
	wireUser(api, middlwareAuth, repo, logger, config)
	wireAuth(api, middlwareAuth, repo, logger, config)
	return router
}

func wireUser(router *gin.RouterGroup, middlwareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseUser := usecase.NewUserService(repo, logger, config)
	adaptorUser := adaptor.NewHandlerUser(usecaseUser, logger)
	router.GET("/test-handler", adaptorUser.TestHandler)
}

func wireAuth(router *gin.RouterGroup, middlwareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseAuth := usecase.NewAuthService(repo, logger, config)
	adaptorAuth := adaptor.NewHandlerAuth(usecaseAuth, logger)
	router.POST("/auth/login", adaptorAuth.Login)
	router.POST("/auth/forgot-password", adaptorAuth.ForgotPassword)
	router.POST("/auth/verify-otp", adaptorAuth.VerifyOtp)
	router.POST("/auth/reset-password", adaptorAuth.ResetPassword)
}

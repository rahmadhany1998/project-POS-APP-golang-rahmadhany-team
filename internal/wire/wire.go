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
	wireRevenue(api, middlwareAuth, repo, logger, config)
	wireMenu(api, middlwareAuth, repo, logger, config)
	return router
}

func wireUser(router *gin.RouterGroup, middlewareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseUser := usecase.NewUserService(repo, logger, config)
	adaptorUser := adaptor.NewHandlerUser(usecaseUser, logger)

	// Create an instance of EmailSender from the configuration
	emailSender := utils.NewEmailSender(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPEmail,
		config.SMTPPassword,
	)

	// Use registerAdminRepo from the previously created repo
	usecaseRegisterAdmin := usecase.NewRegisterAdminUsecase(repo.RegisRepo, emailSender)
	adaptorRegisterAdmin := adaptor.NewRegisterAdminHandler(usecaseRegisterAdmin)

	// Route
	router.GET("/user/profile", middlewareAuth.Auth(), adaptorUser.GetProfile)
	router.PUT("/user/profile", middlewareAuth.Auth(), adaptorUser.UpdateProfile)
	router.GET("/admin/list", middlewareAuth.Auth(), adaptorUser.GetAdminList)
	router.PUT("/admin/role", middlewareAuth.Auth(), adaptorUser.UpdateAdminAccess)
	router.GET("/test-handler", adaptorUser.TestHandler)

	// Route register admin
	router.POST("/admin/register", adaptorRegisterAdmin.RegisterAdmin)
}

func wireAuth(router *gin.RouterGroup, middlwareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseAuth := usecase.NewAuthService(repo, logger, config)
	adaptorAuth := adaptor.NewHandlerAuth(usecaseAuth, logger)
	router.POST("/auth/login", adaptorAuth.Login)
	router.POST("/auth/forgot-password", adaptorAuth.ForgotPassword)
	router.POST("/auth/verify-otp", adaptorAuth.VerifyOtp)
	router.POST("/auth/reset-password", adaptorAuth.ResetPassword)
	router.POST("/auth/logout", middlwareAuth.Auth(), adaptorAuth.Logout)
}

func wireRevenue(router *gin.RouterGroup, middlwareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseRevenue := usecase.NewRevenueService(repo, logger, config)
	adaptorRevenue := adaptor.NewHandlerRevenue(usecaseRevenue, logger)
	router.GET("/report/revenue", middlwareAuth.Auth(), adaptorRevenue.GetRevenueSummary)
	router.GET("/report/monthly-revenue", middlwareAuth.Auth(), adaptorRevenue.GetMonthlyRevenue)
	router.GET("/report/top-products", middlwareAuth.Auth(), adaptorRevenue.GetTopProducts)
}

func wireMenu(router *gin.RouterGroup, middlwareAuth middleware.AuthMiddleware, repo repository.Repository, logger *zap.Logger, config utils.Configuration) {
	usecaseMenu := usecase.NewMenuUsecase(repo, logger, config)
	adaptorMenu := adaptor.NewMenuHandler(usecaseMenu)
	menuRoute := router.Group("/menus").Use(middlwareAuth.Auth())

	menuRoute.GET("", adaptorMenu.GetAllMenus)
	menuRoute.GET("/:id", adaptorMenu.GetMenuByID)
	menuRoute.POST("", adaptorMenu.CreateMenu)
	menuRoute.PUT("/:id", adaptorMenu.UpdateMenu)
	menuRoute.DELETE("/:id", adaptorMenu.DeleteMenu)
}

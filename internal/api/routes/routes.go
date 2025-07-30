package routes

import (
    "project-POS-APP-golang-be-team/internal/api/handlers"
    "project-POS-APP-golang-be-team/internal/api/middlewares"
    "project-POS-APP-golang-be-team/internal/api/websocket"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(
    r *gin.Engine,
    authMiddleware middlewares.AuthMiddleware,
    dashboardHandler *handlers.DashboardHandler,
    inventoryHandler *handlers.InventoryHandler,
    inventoryHub *websocket.InventoryHub,
    // ... other handlers
) {
    // Public routes
    public := r.Group("/api/v1")
    {
        // Authentication routes (if any)
        // public.POST("/login", authHandler.Login)
    }

    // Protected routes
    protected := r.Group("/api/v1")
    protected.Use(authMiddleware.AuthRequired())
    {
        // Dashboard routes
        dashboard := protected.Group("/dashboard")
        {
            dashboard.GET("/summary", dashboardHandler.GetDashboardSummary)
            dashboard.GET("/popular-products", dashboardHandler.GetPopularProducts)
            dashboard.GET("/new-products", dashboardHandler.GetNewProducts)
            dashboard.GET("/revenue-overview", dashboardHandler.GetRevenueOverview)
        }

        // Inventory routes
        inventory := protected.Group("/inventory")
        {
            inventory.GET("", inventoryHandler.GetAllProducts)
            inventory.GET("/:id", inventoryHandler.GetProductByID)
            inventory.POST("", inventoryHandler.CreateProduct)
            inventory.PUT("/:id", inventoryHandler.UpdateProduct)
            inventory.DELETE("/:id", inventoryHandler.DeleteProduct)
            inventory.GET("/export", inventoryHandler.ExportInventoryData)
        }

        // WebSocket endpoint
        protected.GET("/ws/inventory", func(c *gin.Context) {
            inventoryHub.HandleWebSocket(c)
        })

        // ... other routes
    }
}
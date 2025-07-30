package handlers

import (
    "net/http"
    "project-POS-APP-golang-be-team/internal/service"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DashboardHandler struct {
    dashboardService service.DashboardService
    logger           *zap.Logger
}

func NewDashboardHandler(dashboardService service.DashboardService, logger *zap.Logger) *DashboardHandler {
    return &DashboardHandler{
        dashboardService: dashboardService,
        logger:           logger,
    }
}

// GetDashboardSummary returns summary data for dashboard
func (h *DashboardHandler) GetDashboardSummary(c *gin.Context) {
    ctx := c.Request.Context()
    summary, err := h.dashboardService.GetDashboardSummary(ctx)
    if err != nil {
        h.logger.Error("Failed to get dashboard summary", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dashboard summary"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Dashboard summary retrieved successfully",
        "data": gin.H{
            "daily_sales":   summary.DailySales,
            "monthly_sales": summary.MonthlySales,
            "total_revenue": summary.TotalRevenue,
            "table_count":   summary.TableCount,
        },
    })
}

// GetPopularProducts returns list of popular products
func (h *DashboardHandler) GetPopularProducts(c *gin.Context) {
    ctx := c.Request.Context()
    products, err := h.dashboardService.GetPopularProducts(ctx)
    if err != nil {
        h.logger.Error("Failed to get popular products", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get popular products"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Popular products retrieved successfully",
        "data":    products,
    })
}

// GetNewProducts returns list of new products
func (h *DashboardHandler) GetNewProducts(c *gin.Context) {
    ctx := c.Request.Context()
    products, err := h.dashboardService.GetNewProducts(ctx)
    if err != nil {
        h.logger.Error("Failed to get new products", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get new products"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "New products retrieved successfully",
        "data":    products,
    })
}

// GetRevenueOverview returns monthly revenue data for the chart
func (h *DashboardHandler) GetRevenueOverview(c *gin.Context) {
    ctx := c.Request.Context()
    overview, err := h.dashboardService.GetRevenueOverview(ctx)
    if err != nil {
        h.logger.Error("Failed to get revenue overview", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get revenue overview"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Revenue overview retrieved successfully",
        "data":    overview,
    })
}
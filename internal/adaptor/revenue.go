package adaptor

import (
	"net/http"
	"project-POS-APP-golang-be-team/internal/usecase"
	"project-POS-APP-golang-be-team/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerRevenue struct {
	RevenueService usecase.RevenueService
	Logger         *zap.Logger
}

func NewHandlerRevenue(revenue usecase.RevenueService, logger *zap.Logger) HandlerRevenue {
	return HandlerRevenue{
		RevenueService: revenue,
		Logger:         logger,
	}
}

func (h *HandlerRevenue) GetRevenueSummary(ctx *gin.Context) {
	start := ctx.Query("start")
	end := ctx.Query("end")

	report, err := h.RevenueService.GetRevenueSummary(ctx.Request.Context(), start, end)
	if err != nil {
		response.ResponseBadRequest(ctx, http.StatusInternalServerError, "failed to get revenue summary")
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "success", report)
}

func (h *HandlerRevenue) GetMonthlyRevenue(ctx *gin.Context) {
	start := ctx.Query("start")
	end := ctx.Query("end")

	monthly, err := h.RevenueService.GetMonthlyRevenue(ctx.Request.Context(), start, end)
	if err != nil {
		response.ResponseBadRequest(ctx, http.StatusInternalServerError, "failed to get monthly revenue")
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "success", monthly)
}

func (h *HandlerRevenue) GetTopProducts(ctx *gin.Context) {
	start := ctx.Query("start")
	end := ctx.Query("end")

	topProducts, err := h.RevenueService.GetTopProducts(ctx.Request.Context(), start, end)
	if err != nil {
		response.ResponseBadRequest(ctx, http.StatusInternalServerError, "failed to get top products")
		return
	}
	response.ResponseSuccess(ctx, http.StatusOK, "success", topProducts)
}

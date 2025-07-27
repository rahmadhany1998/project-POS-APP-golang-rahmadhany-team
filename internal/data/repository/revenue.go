package repository

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/dto"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RevenueRepository interface {
	GetRevenueReport(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error)
}

type revenueRepositoryImpl struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewRevenueRepository(db *gorm.DB, log *zap.Logger) RevenueRepository {
	return &revenueRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *revenueRepositoryImpl) GetRevenueReport(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error) {
	var report dto.RevenueReport

	// Total revenue
	err := r.DB.WithContext(ctx).
		Model(&entity.Order{}).
		Select("SUM(total) as total").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Scan(&report).Error
	if err != nil {
		r.Log.Error("failed to calculate total revenue", zap.String("error", err.Error()))
		return nil, err
	}

	// Breakdown by status
	statusBreakdown := []struct {
		Status string `json:"status"`
		Count  int    `json:"count"`
	}{}
	err = r.DB.WithContext(ctx).
		Model(&entity.Order{}).
		Select("status, COUNT(*) as count").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Group("status").
		Scan(&statusBreakdown).Error
	if err != nil {
		r.Log.Error("failed to get status breakdown", zap.String("error", err.Error()))
		return nil, err
	}
	report.StatusBreakdown = make(map[string]int)
	for _, s := range statusBreakdown {
		report.StatusBreakdown[s.Status] = s.Count
	}

	// Monthly revenue
	err = r.DB.WithContext(ctx).
		Raw(`
			SELECT TO_CHAR(created_at, 'YYYY-MM') AS month, SUM(total) AS total
			FROM orders
			WHERE created_at BETWEEN ? AND ?
			GROUP BY month
			ORDER BY month ASC
		`, startDate, endDate).Scan(&report.MonthlyRevenue).Error
	if err != nil {
		r.Log.Error("failed to get monthly revenue", zap.String("error", err.Error()))
		return nil, err
	}

	// Top products
	topProducts := []struct {
		Name         string  `json:"name"`
		SellPrice    float64 `json:"sell_price"`
		TotalRevenue float64 `json:"total_revenue"`
		RevenueDate  string  `json:"revenue_date"`
	}{}

	err = r.DB.WithContext(ctx).
		Raw(`
			SELECT p.name, p.price AS sell_price,
				sum(oi.price * oi.quantity) as total_revenue,
				DATE(o.created_at) as revenue_date
			FROM order_items oi
			JOIN products p ON p.id = oi.product_id
			JOIN orders o ON o.id = oi.order_id
			WHERE o.created_at BETWEEN ? AND ?
			GROUP BY p.name, p.price, revenue_date
			ORDER BY total_revenue DESC
			LIMIT 10
		`, startDate, endDate).Scan(&topProducts).Error
	if err != nil {
		r.Log.Error("failed to get top products", zap.String("error", err.Error()))
		return nil, err
	}

	for _, p := range topProducts {
		report.TopProducts = append(report.TopProducts, dto.TopProduct{
			Name:         p.Name,
			SellPrice:    p.SellPrice,
			TotalRevenue: p.TotalRevenue,
			RevenueDate:  p.RevenueDate,
		})
	}

	return &report, nil
}

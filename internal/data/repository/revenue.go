package repository

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/dto"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RevenueRepository interface {
	GetRevenueSummary(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error)
	GetMonthlyRevenue(ctx context.Context, startDate, endDate string) ([]dto.MonthlyRevenue, error)
	GetTopProducts(ctx context.Context, startDate, endDate string) ([]dto.TopProduct, error)
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

func (r *revenueRepositoryImpl) GetRevenueSummary(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error) {
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

	return &report, nil
}

func (r *revenueRepositoryImpl) GetMonthlyRevenue(ctx context.Context, startDate, endDate string) ([]dto.MonthlyRevenue, error) {
	type result struct {
		Month time.Time
		Total float64
	}

	var rawResult []result
	err := r.DB.WithContext(ctx).
		Raw(`
			SELECT DATE_TRUNC('month', created_at) AS month, SUM(total) AS total
			FROM orders
			WHERE created_at BETWEEN ? AND ?
			GROUP BY month
			ORDER BY month ASC
		`, startDate, endDate).
		Scan(&rawResult).Error
	if err != nil {
		r.Log.Error("failed to get monthly revenue", zap.String("error", err.Error()))
		return nil, err
	}

	// Fill all 12 months with default zero if missing
	monthlyMap := make(map[string]float64)
	for _, row := range rawResult {
		monthlyMap[row.Month.Format("January")] = row.Total
	}

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var monthly []dto.MonthlyRevenue
	for _, month := range months {
		monthly = append(monthly, dto.MonthlyRevenue{
			Month: month,
			Total: monthlyMap[month],
		})
	}

	return monthly, nil
}

func (r *revenueRepositoryImpl) GetTopProducts(ctx context.Context, startDate, endDate string) ([]dto.TopProduct, error) {
	var result []dto.TopProduct

	err := r.DB.WithContext(ctx).
		Raw(`
			SELECT p.name, p.price AS sell_price,
				sum(oi.price * oi.quantity) as total_revenue,
				DATE(o.created_at) as revenue_date
			FROM order_items oi
			JOIN products p ON p.id = oi.product_id
			JOIN orders o ON o.id = oi.order_id
			WHERE o.created_at BETWEEN ? AND ?
			GROUP BY p.name, p.price, DATE(o.created_at)
			ORDER BY total_revenue DESC
			LIMIT 10
		`, startDate, endDate).
		Scan(&result).Error
	if err != nil {
		r.Log.Error("failed to get top products", zap.String("error", err.Error()))
		return nil, err
	}
	return result, nil
}

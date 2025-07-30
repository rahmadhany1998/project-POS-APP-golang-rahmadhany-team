package usecase

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"

	"go.uber.org/zap"
)

type RevenueService interface {
	GetRevenueSummary(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error)
	GetMonthlyRevenue(ctx context.Context, startDate, endDate string) ([]dto.MonthlyRevenue, error)
	GetTopProducts(ctx context.Context, startDate, endDate string) ([]dto.TopProduct, error)
}

type revenueService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewRevenueService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) RevenueService {
	return &revenueService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *revenueService) GetRevenueSummary(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error) {
	report, err := s.Repo.RevenueRepo.GetRevenueSummary(ctx, startDate, endDate)
	if err != nil {
		s.Logger.Error("failed to get revenue summary", zap.String("error", err.Error()))
		return nil, err
	}
	return report, nil
}

func (s *revenueService) GetMonthlyRevenue(ctx context.Context, startDate, endDate string) ([]dto.MonthlyRevenue, error) {
	return s.Repo.RevenueRepo.GetMonthlyRevenue(ctx, startDate, endDate)
}

func (s *revenueService) GetTopProducts(ctx context.Context, startDate, endDate string) ([]dto.TopProduct, error) {
	products, err := s.Repo.RevenueRepo.GetTopProducts(ctx, startDate, endDate)
	if err != nil {
		s.Logger.Error("failed to get top products", zap.String("error", err.Error()))
		return nil, err
	}

	for i := range products {
		revenue := products[i].TotalRevenue
		products[i].Margin = s.Config.Margin
		products[i].Profit = products[i].SellPrice * products[i].Margin * (revenue / products[i].SellPrice)
	}

	return products, nil
}

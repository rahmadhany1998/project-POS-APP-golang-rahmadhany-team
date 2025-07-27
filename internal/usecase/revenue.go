package usecase

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"

	"go.uber.org/zap"
)

type RevenueService interface {
	GetRevenueReport(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error)
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

func (s *revenueService) GetRevenueReport(ctx context.Context, startDate, endDate string) (*dto.RevenueReport, error) {
	report, err := s.Repo.RevenueRepo.GetRevenueReport(ctx, startDate, endDate)
	if err != nil {
		s.Logger.Error("failed to get revenue report", zap.String("error", err.Error()))
		return nil, err
	}

	// Inject profit and margin (15%) manually
	for i := range report.TopProducts {
		revenue := report.TopProducts[i].TotalRevenue
		report.TopProducts[i].Margin = s.Config.Margin
		report.TopProducts[i].Profit = report.TopProducts[i].SellPrice * report.TopProducts[i].Margin * (revenue / report.TopProducts[i].SellPrice)
	}

	return report, nil
}

/project-POS-APP-golang-rahmadhany-team/internal/service/dashboard_service.go
package service

import (
    "context"
    "project-POS-APP-golang-be-team/internal/data/entity"
    "project-POS-APP-golang-be-team/internal/data/repository"
    "time"
)

type DashboardService interface {
    GetDashboardSummary(ctx context.Context) (*DashboardSummary, error)
    GetRevenueOverview(ctx context.Context) (*RevenueOverview, error)
    GetPopularProducts(ctx context.Context) ([]ProductResponse, error)
    GetNewProducts(ctx context.Context) ([]ProductResponse, error)
}

type dashboardService struct {
    dashboardRepo repository.DashboardRepository
    productRepo   repository.ProductRepository
}

type DashboardSummary struct {
    DailySales     int `json:"daily_sales"`
    MonthlySales   int `json:"monthly_sales"`
    TotalRevenue   int `json:"total_revenue"`
    TableCount     int `json:"table_count"`
    TotalProducts  int `json:"total_products"`
    TotalUsers     int `json:"total_users"`
}

type RevenueOverview struct {
    MonthlyData map[string]int `json:"monthly_data"` 
}

type ProductResponse struct {
    ID         uint   `json:"id"`
    Name       string `json:"name"`
    Photo      string `json:"photo"`
    Price      int    `json:"price"`
    Stock      int    `json:"stock"`
    Category   string `json:"category"`
    InStock    bool   `json:"in_stock"`
    OrderCount int    `json:"order_count,omitempty"`
}

func NewDashboardService(dashboardRepo repository.DashboardRepository, productRepo repository.ProductRepository) DashboardService {
    return &dashboardService{
        dashboardRepo: dashboardRepo,
        productRepo:   productRepo,
    }
}

func (s *dashboardService) GetDashboardSummary(ctx context.Context) (*DashboardSummary, error) {
    dashboard, err := s.dashboardRepo.GetDashboardSummary(ctx)
    if err != nil {
        return nil, err
    }
    
    dailySales, err := s.dashboardRepo.GetDailySales(ctx)
    if err != nil {
        return nil, err
    }
    
    monthlySales, err := s.dashboardRepo.GetMonthlySales(ctx)
    if err != nil {
        return nil, err
    }
    
    tableCount, err := s.dashboardRepo.GetTableCount(ctx)
    if err != nil {
        return nil, err
    }

    return &DashboardSummary{
        DailySales:    dailySales,
        MonthlySales:  monthlySales,
        TotalRevenue:  dashboard.TotalRevenue,
        TableCount:    tableCount,
        TotalProducts: dashboard.TotalProducts,
        TotalUsers:    dashboard.TotalUsers,
    }, nil
}

func (s *dashboardService) GetRevenueOverview(ctx context.Context) (*RevenueOverview, error) {
    currentYear := time.Now().Year()
    monthlyData, err := s.dashboardRepo.GetMonthlyRevenueData(ctx, currentYear)
    if err != nil {
        return nil, err
    }

    return &RevenueOverview{
        MonthlyData: monthlyData,
    }, nil
}

func (s *dashboardService) GetPopularProducts(ctx context.Context) ([]ProductResponse, error) {
    products, err := s.dashboardRepo.GetPopularProducts(ctx, 5)
    if err != nil {
        return nil, err
    }

    return mapProductsToResponse(products), nil
}

func (s *dashboardService) GetNewProducts(ctx context.Context) ([]ProductResponse, error) {
    products, err := s.dashboardRepo.GetNewProducts(ctx, 5)
    if err != nil {
        return nil, err
    }

    return mapProductsToResponse(products), nil
}

func mapProductsToResponse(products []entity.Product) []ProductResponse {
    result := make([]ProductResponse, 0, len(products))
    
    for _, p := range products {
        result = append(result, ProductResponse{
            ID:       p.ID,
            Name:     p.Name,
            Photo:    p.Photo,
            Price:    p.Price,
            Stock:    p.Stock,
            Category: p.Category.Name,
            InStock:  p.Stock > 0 && p.Available,
        })
    }
    
    return result
}
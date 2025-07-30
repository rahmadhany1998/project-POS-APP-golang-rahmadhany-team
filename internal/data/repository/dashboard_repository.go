package repository

import (
    "context"
    "project-POS-APP-golang-be-team/internal/data/entity"
    "time"

    "gorm.io/gorm"
)

type DashboardRepository interface {
    GetDashboardSummary(ctx context.Context) (*entity.Dashboard, error)
    GetDailySales(ctx context.Context) (int, error)
    GetMonthlySales(ctx context.Context) (int, error)
    GetTableCount(ctx context.Context) (int, error)
    GetPopularProducts(ctx context.Context, limit int) ([]entity.Product, error)
    GetNewProducts(ctx context.Context, limit int) ([]entity.Product, error)
    GetMonthlyRevenueData(ctx context.Context, year int) (map[string]int, error)
}

type dashboardRepository struct {
    db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
    return &dashboardRepository{
        db: db,
    }
}

func (r *dashboardRepository) GetDashboardSummary(ctx context.Context) (*entity.Dashboard, error) {
    var dashboard entity.Dashboard

    // Count total orders
    if err := r.db.Model(&entity.Order{}).Count(&dashboard.TotalOrders).Error; err != nil {
        return nil, err
    }

    // Count total products
    if err := r.db.Model(&entity.Product{}).Count(&dashboard.TotalProducts).Error; err != nil {
        return nil, err
    }

    // Count total users
    if err := r.db.Model(&entity.User{}).Count(&dashboard.TotalUsers).Error; err != nil {
        return nil, err
    }

    // Count total tables
    if err := r.db.Model(&entity.Table{}).Count(&dashboard.TotalTables).Error; err != nil {
        return nil, err
    }

    // Count total revenue (sum of all order totals)
    if err := r.db.Model(&entity.Order{}).Select("COALESCE(SUM(total), 0) as total_revenue").Scan(&dashboard.TotalRevenue).Error; err != nil {
        return nil, err
    }

    // Add other aggregations as needed

    return &dashboard, nil
}

func (r *dashboardRepository) GetDailySales(ctx context.Context) (int, error) {
    var result int
    today := time.Now().Format("2006-01-02")
    
    if err := r.db.Model(&entity.Order{}).
        Where("DATE(created_at) = ?", today).
        Count(&result).Error; err != nil {
        return 0, err
    }
    
    return result, nil
}

func (r *dashboardRepository) GetMonthlySales(ctx context.Context) (int, error) {
    var result int
    currentYear, currentMonth, _ := time.Now().Date()
    
    if err := r.db.Model(&entity.Order{}).
        Where("EXTRACT(YEAR FROM created_at) = ? AND EXTRACT(MONTH FROM created_at) = ?", 
            currentYear, currentMonth).
        Count(&result).Error; err != nil {
        return 0, err
    }
    
    return result, nil
}

func (r *dashboardRepository) GetTableCount(ctx context.Context) (int, error) {
    var count int64
    if err := r.db.Model(&entity.Table{}).Count(&count).Error; err != nil {
        return 0, err
    }
    return int(count), nil
}

func (r *dashboardRepository) GetPopularProducts(ctx context.Context, limit int) ([]entity.Product, error) {
    var products []entity.Product
    
    if err := r.db.Model(&entity.OrderItem{}).
        Select("product_id, COUNT(*) as order_count").
        Group("product_id").
        Order("order_count DESC").
        Limit(limit).
        Scan(&products).Error; err != nil {
        return nil, err
    }
    
    // Fetch full product details
    var productIDs []uint
    for _, p := range products {
        productIDs = append(productIDs, p.ID)
    }
    
    if err := r.db.Preload("Category").
        Where("id IN ?", productIDs).
        Find(&products).Error; err != nil {
        return nil, err
    }
    
    return products, nil
}

func (r *dashboardRepository) GetNewProducts(ctx context.Context, limit int) ([]entity.Product, error) {
    var products []entity.Product
    thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
    
    if err := r.db.Preload("Category").
        Where("created_at >= ?", thirtyDaysAgo).
        Order("created_at DESC").
        Limit(limit).
        Find(&products).Error; err != nil {
        return nil, err
    }
    
    return products, nil
}

func (r *dashboardRepository) GetMonthlyRevenueData(ctx context.Context, year int) (map[string]int, error) {
    type MonthlyRevenue struct {
        Month   int `json:"month"`
        Revenue int `json:"revenue"`
    }
    
    var monthlyData []MonthlyRevenue
    result := make(map[string]int)
    
    if err := r.db.Model(&entity.Order{}).
        Select("EXTRACT(MONTH FROM created_at) as month, COALESCE(SUM(total), 0) as revenue").
        Where("EXTRACT(YEAR FROM created_at) = ?", year).
        Group("month").
        Order("month").
        Scan(&monthlyData).Error; err != nil {
        return nil, err
    }
    
    // Map months to their names and populate result
    months := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
    for _, data := range monthlyData {
        if data.Month >= 1 && data.Month <= 12 {
            result[months[data.Month-1]] = data.Revenue
        }
    }
    
    return result, nil
}
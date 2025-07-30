package repository

import (
    "context"
    "project-POS-APP-golang-be-team/internal/data/entity"
    "time"

    "gorm.io/gorm"
)

type ProductRepository interface {
    FindAll(ctx context.Context, page, limit int, filter string) ([]entity.Product, int64, error)
    FindByID(ctx context.Context, id uint) (*entity.Product, error)
    Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
    Update(ctx context.Context, product *entity.Product) (*entity.Product, error)
    Delete(ctx context.Context, id uint) error
    GetProductExportData(ctx context.Context, startDate, endDate time.Time) ([]entity.Product, int, int, int, error)
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{
        db: db,
    }
}

func (r *productRepository) FindAll(ctx context.Context, page, limit int, filter string) ([]entity.Product, int64, error) {
    var products []entity.Product
    var total int64
    
    offset := (page - 1) * limit
    query := r.db.Model(&entity.Product{}).Preload("Category")
    
    if filter != "" {
        query = query.Where("name LIKE ? OR item_code LIKE ?", "%"+filter+"%", "%"+filter+"%")
    }
    
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
        return nil, 0, err
    }
    
    return products, total, nil
}

func (r *productRepository) FindByID(ctx context.Context, id uint) (*entity.Product, error) {
    var product entity.Product
    
    if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
        return nil, err
    }
    
    return &product, nil
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
    if err := r.db.Create(product).Error; err != nil {
        return nil, err
    }
    
    // Reload with category
    if err := r.db.Preload("Category").First(product, product.ID).Error; err != nil {
        return nil, err
    }
    
    return product, nil
}

func (r *productRepository) Update(ctx context.Context, product *entity.Product) (*entity.Product, error) {
    if err := r.db.Save(product).Error; err != nil {
        return nil, err
    }
    
    // Reload with category
    if err := r.db.Preload("Category").First(product, product.ID).Error; err != nil {
        return nil, err
    }
    
    return product, nil
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
    return r.db.Delete(&entity.Product{}, id).Error
}

func (r *productRepository) GetProductExportData(ctx context.Context, startDate, endDate time.Time) ([]entity.Product, int, int, int, error) {
    var products []entity.Product
    var totalOrders, totalSales, totalRevenue int
    
    // Get products with order counts and revenue in the date range
    // This is a simplified query, in a real application you'd join with order items
    err := r.db.Model(&entity.Product{}).
        Joins("LEFT JOIN order_items ON order_items.product_id = products.id").
        Joins("LEFT JOIN orders ON orders.id = order_items.order_id").
        Where("orders.created_at BETWEEN ? AND ?", startDate, endDate).
        Group("products.id").
        Select("products.*, COUNT(order_items.id) as quantity, SUM(order_items.price * order_items.quantity) as price").
        Preload("Category").
        Find(&products).Error
    
    if err != nil {
        return nil, 0, 0, 0, err
    }
    
    // Get summary data
    type SummaryData struct {
        TotalOrders  int
        TotalSales   int
        TotalRevenue int
    }
    
    var summary SummaryData
    err = r.db.Model(&entity.Order{}).
        Where("created_at BETWEEN ? AND ?", startDate, endDate).
        Select("COUNT(*) as total_orders, SUM(total_items) as total_sales, SUM(total) as total_revenue").
        Scan(&summary).Error
    
    if err != nil {
        return nil, 0, 0, 0, err
    }
    
    return products, summary.TotalOrders, summary.TotalSales, summary.TotalRevenue, nil
}
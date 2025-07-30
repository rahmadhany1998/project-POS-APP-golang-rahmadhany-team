package service

import (
    "context"
    "project-POS-APP-golang-be-team/internal/data/entity"
    "project-POS-APP-golang-be-team/internal/data/repository"
    "time"
)

type InventoryService interface {
    GetAllProducts(ctx context.Context, page, limit int, filter string) ([]ProductResponse, int64, error)
    GetProductByID(ctx context.Context, id uint) (*ProductResponse, error)
    CreateProduct(ctx context.Context, request CreateProductRequest) (*ProductResponse, error)
    UpdateProduct(ctx context.Context, id uint, request UpdateProductRequest) (*ProductResponse, error)
    DeleteProduct(ctx context.Context, id uint) error
    ExportInventoryData(ctx context.Context, month int, year int) (*ExportData, error)
}

type inventoryService struct {
    productRepo repository.ProductRepository
    categoryRepo repository.CategoryRepository
}

type ExportData struct {
    Month       string `json:"month"`
    Year        int    `json:"year"`
    TotalOrders int    `json:"total_orders"`
    TotalSales  int    `json:"total_sales"`
    Revenue     int    `json:"revenue"`
    Products    []ProductExportData `json:"products"`
}

type ProductExportData struct {
    Name       string `json:"name"`
    OrderCount int    `json:"order_count"`
    Revenue    int    `json:"revenue"`
}

type CreateProductRequest struct {
    Name       string `json:"name" binding:"required"`
    Photo      string `json:"photo"`
    ItemCode   string `json:"item_code" binding:"required"`
    Stock      int    `json:"stock"`
    CategoryID uint   `json:"category_id" binding:"required"`
    Price      int    `json:"price" binding:"required"`
    Available  bool   `json:"available"`
    Quantity   int    `json:"quantity"`
    Unit       string `json:"unit"`
}

type UpdateProductRequest struct {
    Name       *string `json:"name"`
    Photo      *string `json:"photo"`
    ItemCode   *string `json:"item_code"`
    Stock      *int    `json:"stock"`
    CategoryID *uint   `json:"category_id"`
    Price      *int    `json:"price"`
    Available  *bool   `json:"available"`
    Quantity   *int    `json:"quantity"`
    Unit       *string `json:"unit"`
}

func NewInventoryService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) InventoryService {
    return &inventoryService{
        productRepo: productRepo,
        categoryRepo: categoryRepo,
    }
}

func (s *inventoryService) GetAllProducts(ctx context.Context, page, limit int, filter string) ([]ProductResponse, int64, error) {
    products, total, err := s.productRepo.FindAll(ctx, page, limit, filter)
    if err != nil {
        return nil, 0, err
    }
    
    return mapProductsToResponse(products), total, nil
}

func (s *inventoryService) GetProductByID(ctx context.Context, id uint) (*ProductResponse, error) {
    product, err := s.productRepo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    response := mapProductToResponse(*product)
    return &response, nil
}

func (s *inventoryService) CreateProduct(ctx context.Context, request CreateProductRequest) (*ProductResponse, error) {
    // Validate category exists
    _, err := s.categoryRepo.FindByID(ctx, request.CategoryID)
    if err != nil {
        return nil, err
    }
    
    product := entity.Product{
        Name:       request.Name,
        Photo:      request.Photo,
        ItemCode:   request.ItemCode,
        Stock:      request.Stock,
        CategoryID: request.CategoryID,
        Price:      request.Price,
        Available:  request.Available,
        Quantity:   request.Quantity,
        Unit:       request.Unit,
        Status:     getStockStatus(request.Stock),
    }
    
    createdProduct, err := s.productRepo.Create(ctx, &product)
    if err != nil {
        return nil, err
    }
    
    response := mapProductToResponse(*createdProduct)
    return &response, nil
}

func (s *inventoryService) UpdateProduct(ctx context.Context, id uint, request UpdateProductRequest) (*ProductResponse, error) {
    product, err := s.productRepo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Update fields if provided
    if request.Name != nil {
        product.Name = *request.Name
    }
    
    if request.Photo != nil {
        product.Photo = *request.Photo
    }
    
    if request.ItemCode != nil {
        product.ItemCode = *request.ItemCode
    }
    
    if request.Stock != nil {
        product.Stock = *request.Stock
        product.Status = getStockStatus(*request.Stock)
    }
    
    if request.CategoryID != nil {
        // Validate category exists
        _, err := s.categoryRepo.FindByID(ctx, *request.CategoryID)
        if err != nil {
            return nil, err
        }
        product.CategoryID = *request.CategoryID
    }
    
    if request.Price != nil {
        product.Price = *request.Price
    }
    
    if request.Available != nil {
        product.Available = *request.Available
    }
    
    if request.Quantity != nil {
        product.Quantity = *request.Quantity
    }
    
    if request.Unit != nil {
        product.Unit = *request.Unit
    }
    
    updatedProduct, err := s.productRepo.Update(ctx, product)
    if err != nil {
        return nil, err
    }
    
    response := mapProductToResponse(*updatedProduct)
    return &response, nil
}

func (s *inventoryService) DeleteProduct(ctx context.Context, id uint) error {
    return s.productRepo.Delete(ctx, id)
}

func (s *inventoryService) ExportInventoryData(ctx context.Context, month int, year int) (*ExportData, error) {
    startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
    endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)
    
    // Get orders for the selected month
    products, totalOrders, totalSales, totalRevenue, err := s.productRepo.GetProductExportData(ctx, startDate, endDate)
    if err != nil {
        return nil, err
    }
    
    // Format month name
    monthName := startDate.Format("January")
    
    // Map products to export data format
    productExport := make([]ProductExportData, 0, len(products))
    for _, p := range products {
        productExport = append(productExport, ProductExportData{
            Name:       p.Name,
            OrderCount: p.Quantity, // Using Quantity field to represent order count
            Revenue:    p.Price,    // Using Price field to represent revenue for this product
        })
    }
    
    return &ExportData{
        Month:       monthName,
        Year:        year,
        TotalOrders: totalOrders,
        TotalSales:  totalSales,
        Revenue:     totalRevenue,
        Products:    productExport,
    }, nil
}

func getStockStatus(stock int) string {
    if stock > 0 {
        return "in stock"
    }
    return "out of stock"
}

func mapProductToResponse(product entity.Product) ProductResponse {
    return ProductResponse{
        ID:       product.ID,
        Name:     product.Name,
        Photo:    product.Photo,
        Price:    product.Price,
        Stock:    product.Stock,
        Category: product.Category.Name,
        InStock:  product.Stock > 0 && product.Available,
    }
}
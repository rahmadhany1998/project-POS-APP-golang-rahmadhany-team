package handlers

import (
    "fmt"
    "net/http"
    "project-POS-APP-golang-be-team/internal/service"
    "strconv"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type InventoryHandler struct {
    inventoryService service.InventoryService
    logger           *zap.Logger
}

func NewInventoryHandler(inventoryService service.InventoryService, logger *zap.Logger) *InventoryHandler {
    return &InventoryHandler{
        inventoryService: inventoryService,
        logger:           logger,
    }
}

// GetAllProducts returns paginated list of products with optional filtering
func (h *InventoryHandler) GetAllProducts(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")
    filter := c.DefaultQuery("filter", "")
    
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }
    
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit < 1 {
        limit = 10
    }
    
    ctx := c.Request.Context()
    products, total, err := h.inventoryService.GetAllProducts(ctx, page, limit, filter)
    if err != nil {
        h.logger.Error("Failed to get products", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Products retrieved successfully",
        "data":    products,
        "meta": gin.H{
            "page":  page,
            "limit": limit,
            "total": total,
        },
    })
}

// GetProductByID returns a single product by ID
func (h *InventoryHandler) GetProductByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }
    
    ctx := c.Request.Context()
    product, err := h.inventoryService.GetProductByID(ctx, uint(id))
    if err != nil {
        h.logger.Error("Failed to get product", zap.Error(err), zap.String("id", idStr))
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Product retrieved successfully",
        "data":    product,
    })
}

// CreateProduct creates a new product
func (h *InventoryHandler) CreateProduct(c *gin.Context) {
    var request service.CreateProductRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    ctx := c.Request.Context()
    product, err := h.inventoryService.CreateProduct(ctx, request)
    if err != nil {
        h.logger.Error("Failed to create product", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "status":  "success",
        "message": "Product created successfully",
        "data":    product,
    })
}

// UpdateProduct updates an existing product
func (h *InventoryHandler) UpdateProduct(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }
    
    var request service.UpdateProductRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    ctx := c.Request.Context()
    product, err := h.inventoryService.UpdateProduct(ctx, uint(id), request)
    if err != nil {
        h.logger.Error("Failed to update product", zap.Error(err), zap.String("id", idStr))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Product updated successfully",
        "data":    product,
    })
}

// DeleteProduct removes a product
func (h *InventoryHandler) DeleteProduct(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }
    
    ctx := c.Request.Context()
    if err := h.inventoryService.DeleteProduct(ctx, uint(id)); err != nil {
        h.logger.Error("Failed to delete product", zap.Error(err), zap.String("id", idStr))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Product deleted successfully",
    })
}

// ExportInventoryData exports inventory data for a specific month and year
func (h *InventoryHandler) ExportInventoryData(c *gin.Context) {
    monthStr := c.DefaultQuery("month", fmt.Sprintf("%d", int(time.Now().Month())))
    yearStr := c.DefaultQuery("year", fmt.Sprintf("%d", time.Now().Year()))
    
    month, err := strconv.Atoi(monthStr)
    if err != nil || month < 1 || month > 12 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month"})
        return
    }
    
    year, err := strconv.Atoi(yearStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
        return
    }
    
    ctx := c.Request.Context()
    data, err := h.inventoryService.ExportInventoryData(ctx, month, year)
    if err != nil {
        h.logger.Error("Failed to export inventory data", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export inventory data"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Inventory data exported successfully",
        "data":    data,
    })
}
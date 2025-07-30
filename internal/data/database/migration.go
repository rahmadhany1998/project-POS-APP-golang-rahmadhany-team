package database

import (
    "log"
    "project-POS-APP-golang-be-team/internal/data/entity"

    "gorm.io/gorm"
)

// AutoMigrate runs database migration and creates tables based on entities
func AutoMigrate(db *gorm.DB) error {
    log.Println("Running database migrations...")
    
    err := db.AutoMigrate(
        &entity.User{},
        &entity.Category{},
        &entity.Product{},
        &entity.Table{},
        &entity.Order{},
        &entity.OrderItem{},
        &entity.Staff{},
        // Add other entities here
    )
    
    if err != nil {
        log.Printf("Migration failed: %v", err)
        return err
    }
    
    log.Println("Migration completed successfully")
    return nil
}

// SeedInitialData populates database with initial required data
func SeedInitialData(db *gorm.DB) error {
    log.Println("Seeding initial data...")
    
    // Seed categories if none exist
    var categoryCount int64
    db.Model(&entity.Category{}).Count(&categoryCount)
    
    if categoryCount == 0 {
        categories := []entity.Category{
            {Name: "Food", Description: "Food items"},
            {Name: "Beverage", Description: "Drink items"},
            {Name: "Dessert", Description: "Sweet treats"},
            {Name: "Snack", Description: "Light bites"},
        }
        
        for _, category := range categories {
            if err := db.Create(&category).Error; err != nil {
                return err
            }
        }
        log.Println("Categories seeded")
    }
    
    // Seed admin user if none exists
    var userCount int64
    db.Model(&entity.User{}).Count(&userCount)
    
    if userCount == 0 {
        // Create default admin user (with hashed password "admin123")
        adminUser := entity.User{
            Name:     "Administrator",
            Email:    "admin@example.com",
            Password: "$2a$10$7JB720yubVSZvUI0rEqK/.VqGOSsOnQSGG7UBj.LgUAVK6YlWW5Bi", // pre-hashed "admin123"
            Role:     "admin",
        }
        
        if err := db.Create(&adminUser).Error; err != nil {
            return err
        }
        log.Println("Admin user seeded")
    }
    
    // Seed sample products if none exist
    var productCount int64
    db.Model(&entity.Product{}).Count(&productCount)
    
    if productCount == 0 {
        // Get category IDs
        var foodCategory, beverageCategory entity.Category
        db.Where("name = ?", "Food").First(&foodCategory)
        db.Where("name = ?", "Beverage").First(&beverageCategory)
        
        products := []entity.Product{
            {
                Name:        "Chicken Parmesan",
                Photo:       "https://example.com/cdn/chicken-parmesan.jpg",
                ItemCode:    "CP001",
                Stock:       50,
                CategoryID:  foodCategory.ID,
                Price:       12500,
                Available:   true,
                Quantity:    1,
                Unit:        "plate",
                Status:      "in stock",
                RetailPrice: 15000,
            },
            {
                Name:        "Ice Tea",
                Photo:       "https://example.com/cdn/ice-tea.jpg",
                ItemCode:    "BV001",
                Stock:       100,
                CategoryID:  beverageCategory.ID,
                Price:       5000,
                Available:   true,
                Quantity:    1,
                Unit:        "glass",
                Status:      "in stock",
                RetailPrice: 8000,
            },
        }
        
        for _, product := range products {
            if err := db.Create(&product).Error; err != nil {
                return err
            }
        }
        log.Println("Sample products seeded")
    }
    
    log.Println("Data seeding completed successfully")
    return nil
}
package entity

type OrderItem struct {
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     int     `json:"price"`
	Order     Order   `gorm:"foreignKey:OrderID" json:"order"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}

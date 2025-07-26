package entity

type Order struct {
	Model
	Customer   string      `json:"customer"`
	TableID    uint        `json:"table_id"`
	Table      Table       `gorm:"foreignKey:TableID" json:"table"`
	Subtotal   float64     `json:"subtotal"`
	Tax        float64     `json:"tax"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"`
	Method     string      `json:"payment_method"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

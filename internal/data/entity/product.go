package entity

type Product struct {
	Model
	Photo       string   `gorm:"type:text" json:"photo"`
	Name        string   `gorm:"type:varchar(100)" json:"name"`
	ItemCode    string   `gorm:"type:varchar(50);unique" json:"item_code"`
	Stock       int      `gorm:"default:0" json:"stock"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
	Price       int      `json:"price"`
	Available   bool     `json:"available"`
	Quantity    int      `json:"quantity"`
	Unit        string   `json:"unit"`
	Status      string   `json:"status"`
	RetailPrice int      `json:"retail_price"`
}

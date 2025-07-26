package entity

type Category struct {
	Model
	Icon        string `gorm:"type:text" json:"icon"`
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

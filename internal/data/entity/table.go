package entity

type Table struct {
	Model
	Name   string `gorm:"type:varchar(50);unique" json:"name"`
	Status string `gorm:"type:varchar(20)" json:"status"`
}

package entity

type UserAccess struct {
	Model
	UserID uint   `json:"user_id"`
	Module string `gorm:"type:varchar(100)" json:"module"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}

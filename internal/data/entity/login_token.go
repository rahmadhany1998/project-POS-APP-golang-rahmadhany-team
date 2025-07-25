package entity

type LoginToken struct {
	Model
	UserID int    `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Token  string `gorm:"type:varchar(100);uniqueIndex;not null"`
}

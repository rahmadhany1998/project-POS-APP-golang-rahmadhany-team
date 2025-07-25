package entity

import "time"

type PasswordResetToken struct {
	Model
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	OtpCode   string    `gorm:"type:char(4);not null"`
	IsUsed    bool      `gorm:"default:false"`
	ExpiresAt time.Time `gorm:"not null"`
}

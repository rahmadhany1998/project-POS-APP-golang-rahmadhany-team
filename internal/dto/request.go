package dto

import "time"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOtpRequest struct {
	Email string `json:"email" binding:"required,email"`
	Otp   string `json:"otp" binding:"required,len=4"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Otp             string `json:"otp" binding:"required,len=4"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

type UpdateProfileRequest struct {
	Name    string    `json:"name" binding:"required"`
	Email   string    `json:"email" binding:"required,email"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
	DOB     time.Time `json:"dob"`
}

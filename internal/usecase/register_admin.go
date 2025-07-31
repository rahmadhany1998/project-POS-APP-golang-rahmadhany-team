package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"
	"time"
)

type RegisterAdminUsecase interface {
	RegisterAdmin(ctx context.Context, req dto.RegisterAdminRequest) error
}

type registerAdminUsecase struct {
	repo        repository.ResgiterAdminRepository
	emailSender utils.EmailSender
}

func NewRegisterAdminUsecase(repo repository.ResgiterAdminRepository, sender utils.EmailSender) RegisterAdminUsecase {
	return &registerAdminUsecase{repo: repo, emailSender: sender}
}

func generateRandomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (u *registerAdminUsecase) RegisterAdmin(ctx context.Context, req dto.RegisterAdminRequest) error {
	password := generateRandomPassword()
	hashedPassword := utils.HashPassword(password)
	dob, _ := time.Parse("2006-01-02", req.DOB)

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		DOB:      dob,
		Role:     "admin",
		Password: hashedPassword,
	}

	if err := u.repo.CreateAdmin(ctx, user); err != nil {
		return err
	}

	subject := "Akun Admin Telah Dibuat"
	body := fmt.Sprintf("Halo %s,\n\nAkun admin Anda telah dibuat. Berikut adalah password Anda: %s\n\nSilakan login dan segera ubah password Anda.", req.Name, password)

	return u.emailSender.SendEmail(req.Email, subject, body)
}

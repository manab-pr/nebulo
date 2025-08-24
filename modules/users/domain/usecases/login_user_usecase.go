package usecases

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/manab-pr/nebulo/modules/users/domain/constants"
	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"github.com/manab-pr/nebulo/modules/users/domain/repository"
)

type LoginUserUseCase struct {
	userRepo repository.UserRepository
}

func NewLoginUserUseCase(userRepo repository.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, req *entities.LoginRequest) (string, error) {
	// Check if user exists
	_, err := uc.userRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found, please register first")
		}
		return "", err
	}

	// Generate OTP for login
	otp := generateOTP()
	otpExpiry := time.Now().Add(constants.OTPExpirationMinutes * time.Minute)

	err = uc.userRepo.UpdateOTP(ctx, req.PhoneNumber, otp, otpExpiry)
	if err != nil {
		return "", err
	}

	// In production, send OTP via SMS service
	// For now, we'll return it for testing
	return otp, nil
}

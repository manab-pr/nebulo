package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"github.com/manab-pr/nebulo/modules/users/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerifyOTPUseCase struct {
	userRepo repository.UserRepository
}

func NewVerifyOTPUseCase(userRepo repository.UserRepository) *VerifyOTPUseCase {
	return &VerifyOTPUseCase{
		userRepo: userRepo,
	}
}

func (uc *VerifyOTPUseCase) Execute(ctx context.Context, req *entities.VerifyOTPRequest) (*entities.AuthResponse, error) {
	// Get user by phone number
	user, err := uc.userRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if OTP is valid and not expired
	if user.OTP != req.OTP {
		return nil, errors.New("invalid OTP")
	}

	if time.Now().After(user.OTPExpiry) {
		return nil, errors.New("OTP has expired")
	}

	// Clear OTP and mark user as verified
	err = uc.userRepo.ClearOTP(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, expiresAt, err := middleware.GenerateToken(user.ID.Hex(), user.PhoneNumber)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &entities.AuthResponse{
		UserID:      user.ID.Hex(),
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		Token:       token,
		ExpiresAt:   expiresAt,
	}, nil
}
package usecases

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/manab-pr/nebulo/modules/users/domain/constants"
	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"github.com/manab-pr/nebulo/modules/users/domain/repository"
)

type RegisterUserUseCase struct {
	userRepo repository.UserRepository
}

func NewRegisterUserUseCase(userRepo repository.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, req *entities.RegisterRequest) (*entities.User, string, error) {
	// Check if user already exists
	existingUser, err := uc.userRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, "", err
	}

	var user *entities.User
	var otp string

	if existingUser != nil {
		// User exists, generate new OTP
		user = existingUser
		otp = generateOTP()
		otpExpiry := time.Now().Add(constants.OTPExpirationMinutes * time.Minute)

		err = uc.userRepo.UpdateOTP(ctx, req.PhoneNumber, otp, otpExpiry)
		if err != nil {
			return nil, "", err
		}
		user.OTP = otp
		user.OTPExpiry = otpExpiry
	} else {
		// Create new user
		otp = generateOTP()
		otpExpiry := time.Now().Add(constants.OTPExpirationMinutes * time.Minute)

		user = &entities.User{
			PhoneNumber: req.PhoneNumber,
			Name:        req.Name,
			IsVerified:  false,
			OTP:         otp,
			OTPExpiry:   otpExpiry,
		}

		err = uc.userRepo.CreateUser(ctx, user)
		if err != nil {
			return nil, "", err
		}
	}

	// In production, send OTP via SMS service
	// For now, we'll return it in the response for testing
	return user, otp, nil
}

func generateOTP() string {
	maxVal := big.NewInt(constants.OTPMaxValue)
	minVal := big.NewInt(constants.OTPMinValue)

	n, _ := rand.Int(rand.Reader, maxVal.Sub(maxVal, minVal).Add(maxVal, big.NewInt(1)))
	return fmt.Sprintf("%06d", n.Add(n, minVal).Int64())
}

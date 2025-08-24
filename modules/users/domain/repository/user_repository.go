package repository

import (
	"context"
	"time"

	"github.com/manab-pr/nebulo/modules/users/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.User, error)
	GetUserByID(ctx context.Context, userID string) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	UpdateOTP(ctx context.Context, phoneNumber, otp string, expiry time.Time) error
	ClearOTP(ctx context.Context, phoneNumber string) error
}

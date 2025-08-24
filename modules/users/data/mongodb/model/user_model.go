package model

import (
	"time"

	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number"`
	Name        string             `bson:"name"`
	IsVerified  bool               `bson:"is_verified"`
	OTP         string             `bson:"otp,omitempty"`
	OTPExpiry   time.Time          `bson:"otp_expiry,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (m *UserModel) ToEntity() *entities.User {
	return &entities.User{
		ID:          m.ID,
		PhoneNumber: m.PhoneNumber,
		Name:        m.Name,
		IsVerified:  m.IsVerified,
		OTP:         m.OTP,
		OTPExpiry:   m.OTPExpiry,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func FromEntity(user *entities.User) *UserModel {
	return &UserModel{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		IsVerified:  user.IsVerified,
		OTP:         user.OTP,
		OTPExpiry:   user.OTPExpiry,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
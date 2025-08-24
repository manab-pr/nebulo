package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number"`
	Name        string             `bson:"name"`
	IsVerified  bool               `bson:"is_verified"`
	OTP         string             `bson:"otp,omitempty"`
	OTPExpiry   time.Time          `bson:"otp_expiry,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,len=10"`
	Name        string `json:"name" validate:"required,min=2,max=100"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,len=10"`
}

type VerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,len=10"`
	OTP         string `json:"otp" validate:"required,len=6"`
}

type AuthResponse struct {
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Token       string `json:"token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type UserProfile struct {
	ID          string    `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Name        string    `json:"name"`
	IsVerified  bool      `json:"is_verified"`
	CreatedAt   time.Time `json:"created_at"`
}
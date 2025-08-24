package dto

type RegisterResponse struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	OTP     string `json:"otp"` // In production, don't return OTP in response
}

type LoginResponse struct {
	Message string `json:"message"`
	OTP     string `json:"otp"` // In production, don't return OTP in response
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"github.com/manab-pr/nebulo/modules/users/domain/usecases"
	"github.com/manab-pr/nebulo/modules/users/presentation/http/dto"
)

type UserHandler struct {
	registerUseCase       *usecases.RegisterUserUseCase
	loginUseCase          *usecases.LoginUserUseCase
	verifyOTPUseCase      *usecases.VerifyOTPUseCase
	getUserProfileUseCase *usecases.GetUserProfileUseCase
	validator             *validator.Validate
}

func NewUserHandler(
	registerUseCase *usecases.RegisterUserUseCase,
	loginUseCase *usecases.LoginUserUseCase,
	verifyOTPUseCase *usecases.VerifyOTPUseCase,
	getUserProfileUseCase *usecases.GetUserProfileUseCase,
) *UserHandler {
	return &UserHandler{
		registerUseCase:       registerUseCase,
		loginUseCase:          loginUseCase,
		verifyOTPUseCase:      verifyOTPUseCase,
		getUserProfileUseCase: getUserProfileUseCase,
		validator:             validator.New(),
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req entities.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	user, otp, err := h.registerUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to register user",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.RegisterResponse{
		Message: "User registered successfully. OTP sent to your phone.",
		UserID:  user.ID.Hex(),
		OTP:     otp, // Remove this in production
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req entities.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	otp, err := h.loginUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Login failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "OTP sent to your phone.",
		OTP:     otp, // Remove this in production
	})
}

func (h *UserHandler) VerifyOTP(c *gin.Context) {
	var req entities.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	authResponse, err := h.verifyOTPUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "OTP verification failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, authResponse)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "User not authenticated",
		})
		return
	}

	profile, err := h.getUserProfileUseCase.Execute(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to get user profile",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}

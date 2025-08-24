package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/manab-pr/nebulo/modules/devices/domain/entities"
	"github.com/manab-pr/nebulo/modules/devices/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterDeviceUseCase struct {
	deviceRepo repository.DeviceRepository
}

func NewRegisterDeviceUseCase(deviceRepo repository.DeviceRepository) *RegisterDeviceUseCase {
	return &RegisterDeviceUseCase{
		deviceRepo: deviceRepo,
	}
}

func (uc *RegisterDeviceUseCase) Execute(
	ctx context.Context, userID string, req entities.DeviceRegistrationRequest,
) (*entities.Device, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Check if device with same IP already exists for this user
	existingDevice, err := uc.deviceRepo.GetByIPAddress(ctx, userObjectID, req.IPAddress)
	if err == nil && existingDevice != nil {
		return nil, errors.New("device with this IP address already exists in your account")
	}

	// Create new device
	device := &entities.Device{
		ID:               primitive.NewObjectID(),
		UserID:           userObjectID,
		Name:             req.Name,
		IPAddress:        req.IPAddress,
		Type:             req.Type,
		TotalStorage:     req.TotalStorage,
		AvailableStorage: req.TotalStorage,
		UsedStorage:      0,
		Status:           entities.DeviceStatusOnline,
		LastHeartbeat:    time.Now(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	createdDevice, err := uc.deviceRepo.Create(ctx, device)
	if err != nil {
		return nil, err
	}

	return createdDevice, nil
}

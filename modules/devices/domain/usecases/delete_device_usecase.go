package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/devices/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteDeviceUseCase struct {
	deviceRepo repository.DeviceRepository
}

func NewDeleteDeviceUseCase(deviceRepo repository.DeviceRepository) *DeleteDeviceUseCase {
	return &DeleteDeviceUseCase{
		deviceRepo: deviceRepo,
	}
}

func (uc *DeleteDeviceUseCase) Execute(ctx context.Context, userID, deviceID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	deviceObjectID, err := primitive.ObjectIDFromHex(deviceID)
	if err != nil {
		return errors.New("invalid device ID")
	}

	// Check if device exists and belongs to user
	device, err := uc.deviceRepo.GetByID(ctx, userObjectID, deviceObjectID)
	if err != nil || device == nil {
		return errors.New("device not found or does not belong to you")
	}

	// Delete device
	err = uc.deviceRepo.Delete(ctx, userObjectID, deviceObjectID)
	if err != nil {
		return err
	}

	return nil
}

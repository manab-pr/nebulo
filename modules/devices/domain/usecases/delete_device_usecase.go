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

func (uc *DeleteDeviceUseCase) Execute(ctx context.Context, deviceID string) error {
	id, err := primitive.ObjectIDFromHex(deviceID)
	if err != nil {
		return errors.New("invalid device ID")
	}

	// Check if device exists
	device, err := uc.deviceRepo.GetByID(ctx, id)
	if err != nil || device == nil {
		return errors.New("device not found")
	}

	// Delete device
	err = uc.deviceRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
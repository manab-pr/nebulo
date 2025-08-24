package usecases

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/manab-pr/nebulo/modules/devices/domain/entities"
	"github.com/manab-pr/nebulo/modules/devices/domain/repository"
)

type ListDevicesUseCase struct {
	deviceRepo repository.DeviceRepository
}

func NewListDevicesUseCase(deviceRepo repository.DeviceRepository) *ListDevicesUseCase {
	return &ListDevicesUseCase{
		deviceRepo: deviceRepo,
	}
}

func (uc *ListDevicesUseCase) Execute(ctx context.Context, userID string) ([]*entities.Device, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	devices, err := uc.deviceRepo.GetAllByUser(ctx, userObjectID)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (uc *ListDevicesUseCase) GetOnlineDevices(ctx context.Context, userID string) ([]*entities.Device, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	devices, err := uc.deviceRepo.GetOnlineDevicesByUser(ctx, userObjectID)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

package usecases

import (
	"context"
	"errors"

	deviceRepo "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	fileRepo "github.com/manab-pr/nebulo/modules/files/domain/repository"
	"github.com/manab-pr/nebulo/modules/storage/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetDeviceStorageUseCase struct {
	deviceRepo deviceRepo.DeviceRepository
	fileRepo   fileRepo.FileRepository
}

func NewGetDeviceStorageUseCase(deviceRepo deviceRepo.DeviceRepository, fileRepo fileRepo.FileRepository) *GetDeviceStorageUseCase {
	return &GetDeviceStorageUseCase{
		deviceRepo: deviceRepo,
		fileRepo:   fileRepo,
	}
}

func (uc *GetDeviceStorageUseCase) Execute(ctx context.Context, userID, deviceID string) (*entities.DeviceStorageInfo, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	deviceObjectID, err := primitive.ObjectIDFromHex(deviceID)
	if err != nil {
		return nil, errors.New("invalid device ID")
	}

	// Get device (user-scoped)
	device, err := uc.deviceRepo.GetByID(ctx, userObjectID, deviceObjectID)
	if err != nil {
		return nil, err
	}

	if device == nil {
		return nil, errors.New("device not found or does not belong to you")
	}

	// Get files stored on this device (user-scoped)
	files, err := uc.fileRepo.GetByUserAndDeviceID(ctx, userObjectID, deviceObjectID)
	if err != nil {
		return nil, err
	}

	// Build storage info
	storageInfo := &entities.DeviceStorageInfo{
		DeviceID:         device.ID.Hex(),
		DeviceName:       device.Name,
		TotalStorage:     device.TotalStorage,
		UsedStorage:      device.UsedStorage,
		AvailableStorage: device.AvailableStorage,
		FileCount:        len(files),
		Status:           string(device.Status),
	}

	return storageInfo, nil
}

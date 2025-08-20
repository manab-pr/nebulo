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

func (uc *GetDeviceStorageUseCase) Execute(ctx context.Context, deviceID string) (*entities.DeviceStorageInfo, error) {
	id, err := primitive.ObjectIDFromHex(deviceID)
	if err != nil {
		return nil, errors.New("invalid device ID")
	}

	// Get device
	device, err := uc.deviceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if device == nil {
		return nil, errors.New("device not found")
	}

	// Get files stored on this device
	files, err := uc.fileRepo.GetByDeviceID(ctx, id)
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

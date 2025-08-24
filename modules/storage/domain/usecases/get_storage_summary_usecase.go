package usecases

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	deviceEntities "github.com/manab-pr/nebulo/modules/devices/domain/entities"
	deviceRepo "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	fileRepo "github.com/manab-pr/nebulo/modules/files/domain/repository"
	"github.com/manab-pr/nebulo/modules/storage/domain/entities"
)

type GetStorageSummaryUseCase struct {
	deviceRepo deviceRepo.DeviceRepository
	fileRepo   fileRepo.FileRepository
}

func NewGetStorageSummaryUseCase(deviceRepo deviceRepo.DeviceRepository, fileRepo fileRepo.FileRepository) *GetStorageSummaryUseCase {
	return &GetStorageSummaryUseCase{
		deviceRepo: deviceRepo,
		fileRepo:   fileRepo,
	}
}

func (uc *GetStorageSummaryUseCase) Execute(ctx context.Context, userID string) (*entities.StorageSummary, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Get user's devices
	devices, err := uc.deviceRepo.GetAllByUser(ctx, userObjectID)
	if err != nil {
		return nil, err
	}

	// Get user's files
	files, err := uc.fileRepo.GetAllByUser(ctx, userObjectID)
	if err != nil {
		return nil, err
	}

	// Calculate summary
	summary := &entities.StorageSummary{}
	summary.TotalDevices = len(devices)
	summary.TotalFiles = len(files)

	var totalStorage, usedStorage, availableStorage int64
	var onlineCount int

	for _, device := range devices {
		totalStorage += device.TotalStorage
		usedStorage += device.UsedStorage
		availableStorage += device.AvailableStorage

		if device.Status == deviceEntities.DeviceStatusOnline {
			onlineCount++
		}
	}

	summary.OnlineDevices = onlineCount
	summary.OfflineDevices = summary.TotalDevices - onlineCount
	summary.TotalStorage = totalStorage
	summary.UsedStorage = usedStorage
	summary.AvailableStorage = availableStorage

	return summary, nil
}

package usecases

import (
	"context"

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

func (uc *GetStorageSummaryUseCase) Execute(ctx context.Context) (*entities.StorageSummary, error) {
	// Get all devices
	devices, err := uc.deviceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Get all files
	files, err := uc.fileRepo.GetAll(ctx)
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

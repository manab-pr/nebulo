package container

import (
	deviceRepo "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	fileRepo "github.com/manab-pr/nebulo/modules/files/domain/repository"
	storageUseCases "github.com/manab-pr/nebulo/modules/storage/domain/usecases"
	storageHandlers "github.com/manab-pr/nebulo/modules/storage/presentation/http/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)

type StorageContainer struct {
	SummaryUseCase       *storageUseCases.GetStorageSummaryUseCase
	DeviceStorageUseCase *storageUseCases.GetDeviceStorageUseCase
	Handler              *storageHandlers.StorageHandler
}

func NewStorageContainer(db *mongo.Database, deviceRepo deviceRepo.DeviceRepository, fileRepo fileRepo.FileRepository) *StorageContainer {
	// Initialize use cases
	summaryUseCase := storageUseCases.NewGetStorageSummaryUseCase(deviceRepo, fileRepo)
	deviceStorageUseCase := storageUseCases.NewGetDeviceStorageUseCase(deviceRepo, fileRepo)

	// Initialize handler
	handler := storageHandlers.NewStorageHandler(
		summaryUseCase,
		deviceStorageUseCase,
	)

	return &StorageContainer{
		SummaryUseCase:       summaryUseCase,
		DeviceStorageUseCase: deviceStorageUseCase,
		Handler:              handler,
	}
}
package container

import (
	deviceRepo "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	fileRepo "github.com/manab-pr/nebulo/modules/files/data/mongodb/repository"
	fileRepository "github.com/manab-pr/nebulo/modules/files/domain/repository"
	fileUseCases "github.com/manab-pr/nebulo/modules/files/domain/usecases"
	fileHandlers "github.com/manab-pr/nebulo/modules/files/presentation/http/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)

type FileContainer struct {
	Repository    fileRepository.FileRepository
	StoreUseCase  *fileUseCases.StoreFileUseCase
	GetUseCase    *fileUseCases.GetFileUseCase
	DeleteUseCase *fileUseCases.DeleteFileUseCase
	Handler       *fileHandlers.FileHandler
}

func NewFileContainer(db *mongo.Database) *FileContainer {
	// Initialize repository
	repo := fileRepo.NewMongoFileRepository(db)

	// For file operations, we'll need the device repository too
	// We'll pass it from the app container later
	return &FileContainer{
		Repository: repo,
	}
}

func (c *FileContainer) InitializeWithDeviceRepo(deviceRepo deviceRepo.DeviceRepository) {
	// Initialize use cases with dependencies
	storeUseCase := fileUseCases.NewStoreFileUseCase(c.Repository, deviceRepo)
	getUseCase := fileUseCases.NewGetFileUseCase(c.Repository)
	deleteUseCase := fileUseCases.NewDeleteFileUseCase(c.Repository)

	// Initialize handler
	handler := fileHandlers.NewFileHandler(
		storeUseCase,
		getUseCase,
		deleteUseCase,
	)

	c.StoreUseCase = storeUseCase
	c.GetUseCase = getUseCase
	c.DeleteUseCase = deleteUseCase
	c.Handler = handler
}

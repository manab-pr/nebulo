package container

import (
	"github.com/manab-pr/nebulo/config"

	deviceHandlers "github.com/manab-pr/nebulo/modules/devices/presentation/http/handlers"
	fileHandlers "github.com/manab-pr/nebulo/modules/files/presentation/http/handlers"
	searchHandlers "github.com/manab-pr/nebulo/modules/search/presentation/http/handlers"
	storageHandlers "github.com/manab-pr/nebulo/modules/storage/presentation/http/handlers"
	transferHandlers "github.com/manab-pr/nebulo/modules/transfers/presentation/http/handlers"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type AppContainer struct {
	// Configuration and infrastructure
	Config *config.Config
	Logger *zap.Logger
	DB     *mongo.Database
	Redis  *redis.Client

	// Handlers
	DeviceHandler   *deviceHandlers.DeviceHandler
	FileHandler     *fileHandlers.FileHandler
	TransferHandler *transferHandlers.TransferHandler
	StorageHandler  *storageHandlers.StorageHandler
	SearchHandler   *searchHandlers.SearchHandler
}

func NewAppContainer(db *mongo.Database, redis *redis.Client, cfg *config.Config, logger *zap.Logger) *AppContainer {
	container := &AppContainer{
		Config: cfg,
		Logger: logger,
		DB:     db,
		Redis:  redis,
	}

	// Initialize repositories
	deviceContainer := NewDeviceContainer(db)
	fileContainer := NewFileContainer(db)
	fileContainer.InitializeWithDeviceRepo(deviceContainer.Repository)
	transferContainer := NewTransferContainer(db)
	storageContainer := NewStorageContainer(db, deviceContainer.Repository, fileContainer.Repository)
	searchContainer := NewSearchContainer(fileContainer.Repository, deviceContainer.Repository)

	// Set handlers
	container.DeviceHandler = deviceContainer.Handler
	container.FileHandler = fileContainer.Handler
	container.TransferHandler = transferContainer.Handler
	container.StorageHandler = storageContainer.Handler
	container.SearchHandler = searchContainer.Handler

	return container
}

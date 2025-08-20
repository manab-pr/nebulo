package container

import (
	deviceRepo "github.com/manab-pr/nebulo/modules/devices/data/mongodb/repository"
	deviceRepository "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	deviceUseCases "github.com/manab-pr/nebulo/modules/devices/domain/usecases"
	deviceHandlers "github.com/manab-pr/nebulo/modules/devices/presentation/http/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceContainer struct {
	Repository          deviceRepository.DeviceRepository
	RegisterUseCase     *deviceUseCases.RegisterDeviceUseCase
	HeartbeatUseCase    *deviceUseCases.HeartbeatUseCase
	ListDevicesUseCase  *deviceUseCases.ListDevicesUseCase
	DeleteDeviceUseCase *deviceUseCases.DeleteDeviceUseCase
	Handler             *deviceHandlers.DeviceHandler
}

func NewDeviceContainer(db *mongo.Database) *DeviceContainer {
	// Initialize repository
	repo := deviceRepo.NewMongoDeviceRepository(db)

	// Initialize use cases
	registerUseCase := deviceUseCases.NewRegisterDeviceUseCase(repo)
	heartbeatUseCase := deviceUseCases.NewHeartbeatUseCase(repo)
	listDevicesUseCase := deviceUseCases.NewListDevicesUseCase(repo)
	deleteDeviceUseCase := deviceUseCases.NewDeleteDeviceUseCase(repo)

	// Initialize handler
	handler := deviceHandlers.NewDeviceHandler(
		registerUseCase,
		heartbeatUseCase,
		listDevicesUseCase,
		deleteDeviceUseCase,
	)

	return &DeviceContainer{
		Repository:          repo,
		RegisterUseCase:     registerUseCase,
		HeartbeatUseCase:    heartbeatUseCase,
		ListDevicesUseCase:  listDevicesUseCase,
		DeleteDeviceUseCase: deleteDeviceUseCase,
		Handler:             handler,
	}
}

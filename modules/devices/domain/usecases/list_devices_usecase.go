package usecases

import (
	"context"

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

func (uc *ListDevicesUseCase) Execute(ctx context.Context) ([]*entities.Device, error) {
	devices, err := uc.deviceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (uc *ListDevicesUseCase) GetOnlineDevices(ctx context.Context) ([]*entities.Device, error) {
	devices, err := uc.deviceRepo.GetOnlineDevices(ctx)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

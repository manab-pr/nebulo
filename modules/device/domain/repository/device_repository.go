package repository

import (
	"context"

	"github.com/manab-pr/nebulo/modules/device/domain/entities"
)

type DeviceRepository interface {
	RegisterDevice(ctx context.Context, device *entities.Devices) error
	GetDeviceByID(ctx context.Context, id uint64) (*entities.Devices, error)
	GetAllDevices(ctx context.Context, filters map[string]interface{}) ([]entities.Devices, error)
	UpdateDevice(ctx context.Context, id uint64, updates map[string]interface{}) error
	DeleteDevice(ctx context.Context, id uint64) error
}
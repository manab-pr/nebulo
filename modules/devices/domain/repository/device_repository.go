package repository

import (
	"context"
	"github.com/manab-pr/nebulo/modules/devices/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceRepository interface {
	Create(ctx context.Context, device *entities.Device) (*entities.Device, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Device, error)
	GetByIPAddress(ctx context.Context, ipAddress string) (*entities.Device, error)
	GetAll(ctx context.Context) ([]*entities.Device, error)
	GetOnlineDevices(ctx context.Context) ([]*entities.Device, error)
	Update(ctx context.Context, device *entities.Device) error
	UpdateHeartbeat(ctx context.Context, id primitive.ObjectID, availableStorage, usedStorage int64) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status entities.DeviceStatus) error
}
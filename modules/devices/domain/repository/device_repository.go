package repository

import (
	"context"

	"github.com/manab-pr/nebulo/modules/devices/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceRepository interface {
	Create(ctx context.Context, device *entities.Device) (*entities.Device, error)
	GetByID(ctx context.Context, userID, deviceID primitive.ObjectID) (*entities.Device, error)
	GetByIPAddress(ctx context.Context, userID primitive.ObjectID, ipAddress string) (*entities.Device, error)
	GetAllByUser(ctx context.Context, userID primitive.ObjectID) ([]*entities.Device, error)
	GetOnlineDevicesByUser(ctx context.Context, userID primitive.ObjectID) ([]*entities.Device, error)
	Update(ctx context.Context, device *entities.Device) error
	UpdateHeartbeat(ctx context.Context, userID, deviceID primitive.ObjectID, availableStorage, usedStorage int64) error
	Delete(ctx context.Context, userID, deviceID primitive.ObjectID) error
	UpdateStatus(ctx context.Context, userID, deviceID primitive.ObjectID, status entities.DeviceStatus) error
}

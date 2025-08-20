package model

import (
	"time"

	"github.com/manab-pr/nebulo/modules/devices/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceModel struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	IPAddress        string             `bson:"ip_address"`
	Type             string             `bson:"type"`
	TotalStorage     int64              `bson:"total_storage"`
	AvailableStorage int64              `bson:"available_storage"`
	UsedStorage      int64              `bson:"used_storage"`
	Status           string             `bson:"status"`
	LastHeartbeat    time.Time          `bson:"last_heartbeat"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

func (d *DeviceModel) ToEntity() *entities.Device {
	return &entities.Device{
		ID:               d.ID,
		Name:             d.Name,
		IPAddress:        d.IPAddress,
		Type:             d.Type,
		TotalStorage:     d.TotalStorage,
		AvailableStorage: d.AvailableStorage,
		UsedStorage:      d.UsedStorage,
		Status:           entities.DeviceStatus(d.Status),
		LastHeartbeat:    d.LastHeartbeat,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
	}
}

func FromEntity(device *entities.Device) *DeviceModel {
	return &DeviceModel{
		ID:               device.ID,
		Name:             device.Name,
		IPAddress:        device.IPAddress,
		Type:             device.Type,
		TotalStorage:     device.TotalStorage,
		AvailableStorage: device.AvailableStorage,
		UsedStorage:      device.UsedStorage,
		Status:           string(device.Status),
		LastHeartbeat:    device.LastHeartbeat,
		CreatedAt:        device.CreatedAt,
		UpdatedAt:        device.UpdatedAt,
	}
}
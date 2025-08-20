package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	IPAddress        string             `bson:"ip_address"`
	Type             string             `bson:"type"`
	TotalStorage     int64              `bson:"total_storage"`
	AvailableStorage int64              `bson:"available_storage"`
	UsedStorage      int64              `bson:"used_storage"`
	Status           DeviceStatus       `bson:"status"`
	LastHeartbeat    time.Time          `bson:"last_heartbeat"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

type DeviceStatus string

const (
	DeviceStatusOnline  DeviceStatus = "online"
	DeviceStatusOffline DeviceStatus = "offline"
	DeviceStatusFailed  DeviceStatus = "failed"
)

type DeviceRegistrationRequest struct {
	Name         string `json:"name" validate:"required"`
	IPAddress    string `json:"ip_address" validate:"required,ip"`
	Type         string `json:"type" validate:"required"`
	TotalStorage int64  `json:"total_storage" validate:"required,min=1"`
}

type DeviceHeartbeatRequest struct {
	DeviceID         string `json:"device_id" validate:"required"`
	AvailableStorage int64  `json:"available_storage" validate:"min=0"`
	UsedStorage      int64  `json:"used_storage" validate:"min=0"`
}
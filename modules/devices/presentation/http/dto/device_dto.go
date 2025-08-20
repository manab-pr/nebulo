package dto

import (
	"time"

	"github.com/manab-pr/nebulo/modules/devices/domain/entities"
)

type DeviceRegisterRequest struct {
	Name         string `json:"name" validate:"required"`
	IPAddress    string `json:"ip_address" validate:"required"`
	Type         string `json:"type" validate:"required"`
	TotalStorage int64  `json:"total_storage" validate:"required,min=1"`
}

type DeviceHeartbeatRequest struct {
	DeviceID         string `json:"device_id" validate:"required"`
	AvailableStorage int64  `json:"available_storage" validate:"min=0"`
	UsedStorage      int64  `json:"used_storage" validate:"min=0"`
}

type DeviceResponse struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	IPAddress        string    `json:"ip_address"`
	Type             string    `json:"type"`
	TotalStorage     int64     `json:"total_storage"`
	AvailableStorage int64     `json:"available_storage"`
	UsedStorage      int64     `json:"used_storage"`
	Status           string    `json:"status"`
	LastHeartbeat    time.Time `json:"last_heartbeat"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func ToDeviceResponse(device *entities.Device) *DeviceResponse {
	return &DeviceResponse{
		ID:               device.ID.Hex(),
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

func ToDeviceResponses(devices []*entities.Device) []*DeviceResponse {
	responses := make([]*DeviceResponse, len(devices))
	for i, device := range devices {
		responses[i] = ToDeviceResponse(device)
	}
	return responses
}

func (r *DeviceRegisterRequest) ToEntity() entities.DeviceRegistrationRequest {
	return entities.DeviceRegistrationRequest{
		Name:         r.Name,
		IPAddress:    r.IPAddress,
		Type:         r.Type,
		TotalStorage: r.TotalStorage,
	}
}

func (r *DeviceHeartbeatRequest) ToEntity() entities.DeviceHeartbeatRequest {
	return entities.DeviceHeartbeatRequest{
		DeviceID:         r.DeviceID,
		AvailableStorage: r.AvailableStorage,
		UsedStorage:      r.UsedStorage,
	}
}

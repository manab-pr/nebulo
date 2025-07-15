package model

import (
	"time"

	"github.com/manab-pr/nebulo/modules/device/domain/entities"
)

type Device struct {
	ID               uint64    `gorm:"primaryKey" json:"id"`
	UserID           uint64    `gorm:"not null" json:"user_id"`
	Name             string    `gorm:"not null" json:"name"`
	IPAddress        string    `gorm:"not null" json:"ip_address"`
	OS               string    `gorm:"type:text" json:"os"`
	Type             string    `json:"device_type"`
	TotalStorage     int64     `json:"total_storage"`
	AvailableStorage int64     `json:"available_storage"`
	LastSeen         time.Time `json:"last_seen"`
	Status           string    `gorm:"default:active" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FromEntity(e *entities.Devices) *Device {
	return &Device{
		ID:               e.ID,
		UserID:           e.UserID,
		Name:             e.Name,
		IPAddress:        e.IPAddress,
		OS:               e.OS,
		Type:             e.Type,
		TotalStorage:     e.TotalStorage,
		AvailableStorage: e.AvailableStorage,
		LastSeen:         e.LastSeen,
		Status:           e.Status,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
	}
}

func (m *Device) ToEntity() *entities.Devices {
	return &entities.Devices{
		ID:               m.ID,
		UserID:           m.UserID,
		Name:             m.Name,
		IPAddress:        m.IPAddress,
		OS:               m.OS,
		Type:             m.Type,
		TotalStorage:     m.TotalStorage,
		AvailableStorage: m.AvailableStorage,
		LastSeen:         m.LastSeen,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}

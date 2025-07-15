package entities

import "time"

type Devices struct {
	ID               uint64
	UserID           uint64
	Name             string
	IPAddress        string
	OS               string
	Type             string
	TotalStorage     int64
	AvailableStorage int64
	LastSeen         time.Time
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

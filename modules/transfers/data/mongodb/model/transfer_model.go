package model

import (
	"time"

	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FileID      primitive.ObjectID `bson:"file_id"`
	DeviceID    primitive.ObjectID `bson:"device_id"`
	Status      string             `bson:"status"`
	Priority    int                `bson:"priority"`
	Retries     int                `bson:"retries"`
	MaxRetries  int                `bson:"max_retries"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	StartedAt   *time.Time         `bson:"started_at,omitempty"`
	CompletedAt *time.Time         `bson:"completed_at,omitempty"`
	ErrorMsg    string             `bson:"error_msg,omitempty"`
}

func (t *TransferModel) ToEntity() *entities.Transfer {
	return &entities.Transfer{
		ID:          t.ID,
		FileID:      t.FileID,
		DeviceID:    t.DeviceID,
		Status:      entities.TransferStatus(t.Status),
		Priority:    t.Priority,
		Retries:     t.Retries,
		MaxRetries:  t.MaxRetries,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		StartedAt:   t.StartedAt,
		CompletedAt: t.CompletedAt,
		ErrorMsg:    t.ErrorMsg,
	}
}

func FromEntity(transfer *entities.Transfer) *TransferModel {
	return &TransferModel{
		ID:          transfer.ID,
		FileID:      transfer.FileID,
		DeviceID:    transfer.DeviceID,
		Status:      string(transfer.Status),
		Priority:    transfer.Priority,
		Retries:     transfer.Retries,
		MaxRetries:  transfer.MaxRetries,
		CreatedAt:   transfer.CreatedAt,
		UpdatedAt:   transfer.UpdatedAt,
		StartedAt:   transfer.StartedAt,
		CompletedAt: transfer.CompletedAt,
		ErrorMsg:    transfer.ErrorMsg,
	}
}
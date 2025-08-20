package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transfer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FileID     primitive.ObjectID `bson:"file_id"`
	DeviceID   primitive.ObjectID `bson:"device_id"`
	Status     TransferStatus     `bson:"status"`
	Priority   int                `bson:"priority"`
	Retries    int                `bson:"retries"`
	MaxRetries int                `bson:"max_retries"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	StartedAt  *time.Time         `bson:"started_at,omitempty"`
	CompletedAt *time.Time        `bson:"completed_at,omitempty"`
	ErrorMsg   string             `bson:"error_msg,omitempty"`
}

type TransferStatus string

const (
	TransferStatusPending    TransferStatus = "pending"
	TransferStatusInProgress TransferStatus = "in_progress"
	TransferStatusCompleted  TransferStatus = "completed"
	TransferStatusFailed     TransferStatus = "failed"
	TransferStatusCancelled  TransferStatus = "cancelled"
)

type CompleteTransferRequest struct {
	TransferID string `json:"transfer_id" validate:"required"`
	Success    bool   `json:"success"`
	ErrorMsg   string `json:"error_msg,omitempty"`
}

type PendingTransfersResponse struct {
	ID       string `json:"id"`
	FileID   string `json:"file_id"`
	DeviceID string `json:"device_id"`
	Priority int    `json:"priority"`
	Retries  int    `json:"retries"`
}
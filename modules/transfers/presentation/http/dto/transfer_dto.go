package dto

import (
	"time"

	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"
)

type CompleteTransferRequest struct {
	TransferID string `json:"transfer_id" validate:"required"`
	Success    bool   `json:"success"`
	ErrorMsg   string `json:"error_msg,omitempty"`
}

type TransferResponse struct {
	ID          string     `json:"id"`
	FileID      string     `json:"file_id"`
	DeviceID    string     `json:"device_id"`
	Status      string     `json:"status"`
	Priority    int        `json:"priority"`
	Retries     int        `json:"retries"`
	MaxRetries  int        `json:"max_retries"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	ErrorMsg    string     `json:"error_msg,omitempty"`
}

type PendingTransferResponse struct {
	ID       string `json:"id"`
	FileID   string `json:"file_id"`
	Priority int    `json:"priority"`
	Retries  int    `json:"retries"`
}

func ToTransferResponse(transfer *entities.Transfer) *TransferResponse {
	return &TransferResponse{
		ID:          transfer.ID.Hex(),
		FileID:      transfer.FileID.Hex(),
		DeviceID:    transfer.DeviceID.Hex(),
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

func ToPendingTransferResponse(transfer *entities.Transfer) *PendingTransferResponse {
	return &PendingTransferResponse{
		ID:       transfer.ID.Hex(),
		FileID:   transfer.FileID.Hex(),
		Priority: transfer.Priority,
		Retries:  transfer.Retries,
	}
}

func ToTransferResponses(transfers []*entities.Transfer) []*TransferResponse {
	responses := make([]*TransferResponse, len(transfers))
	for i, transfer := range transfers {
		responses[i] = ToTransferResponse(transfer)
	}
	return responses
}

func ToPendingTransferResponses(transfers []*entities.Transfer) []*PendingTransferResponse {
	responses := make([]*PendingTransferResponse, len(transfers))
	for i, transfer := range transfers {
		responses[i] = ToPendingTransferResponse(transfer)
	}
	return responses
}

func (r *CompleteTransferRequest) ToEntity() entities.CompleteTransferRequest {
	return entities.CompleteTransferRequest{
		TransferID: r.TransferID,
		Success:    r.Success,
		ErrorMsg:   r.ErrorMsg,
	}
}

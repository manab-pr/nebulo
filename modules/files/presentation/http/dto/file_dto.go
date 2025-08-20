package dto

import (
	"time"

	"github.com/manab-pr/nebulo/modules/files/domain/entities"
)

type StoreFileRequest struct {
	Name         string `json:"name" validate:"required"`
	Size         int64  `json:"size" validate:"required,min=1"`
	MimeType     string `json:"mime_type"`
	TargetDevice string `json:"target_device,omitempty"`
}

type FileResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	OriginalName string    `json:"original_name"`
	Size         int64     `json:"size"`
	MimeType     string    `json:"mime_type"`
	StoredOn     string    `json:"stored_on"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToFileResponse(file *entities.File) *FileResponse {
	return &FileResponse{
		ID:           file.ID.Hex(),
		Name:         file.Name,
		OriginalName: file.OriginalName,
		Size:         file.Size,
		MimeType:     file.MimeType,
		StoredOn:     file.StoredOn.Hex(),
		Status:       string(file.Status),
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
	}
}

func ToFileResponses(files []*entities.File) []*FileResponse {
	responses := make([]*FileResponse, len(files))
	for i, file := range files {
		responses[i] = ToFileResponse(file)
	}
	return responses
}

func (r *StoreFileRequest) ToEntity() entities.StoreFileRequest {
	return entities.StoreFileRequest{
		Name:         r.Name,
		Size:         r.Size,
		MimeType:     r.MimeType,
		TargetDevice: r.TargetDevice,
	}
}

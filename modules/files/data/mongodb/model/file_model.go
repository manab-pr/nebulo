package model

import (
	"time"

	"github.com/manab-pr/nebulo/modules/files/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	OriginalName string             `bson:"original_name"`
	Size         int64              `bson:"size"`
	MimeType     string             `bson:"mime_type"`
	Checksum     string             `bson:"checksum"`
	StoredOn     primitive.ObjectID `bson:"stored_on"`
	Status       string             `bson:"status"`
	StoragePath  string             `bson:"storage_path"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

func (f *FileModel) ToEntity() *entities.File {
	return &entities.File{
		ID:           f.ID,
		Name:         f.Name,
		OriginalName: f.OriginalName,
		Size:         f.Size,
		MimeType:     f.MimeType,
		Checksum:     f.Checksum,
		StoredOn:     f.StoredOn,
		Status:       entities.FileStatus(f.Status),
		StoragePath:  f.StoragePath,
		CreatedAt:    f.CreatedAt,
		UpdatedAt:    f.UpdatedAt,
	}
}

func FromEntity(file *entities.File) *FileModel {
	return &FileModel{
		ID:           file.ID,
		Name:         file.Name,
		OriginalName: file.OriginalName,
		Size:         file.Size,
		MimeType:     file.MimeType,
		Checksum:     file.Checksum,
		StoredOn:     file.StoredOn,
		Status:       string(file.Status),
		StoragePath:  file.StoragePath,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
	}
}

package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	OriginalName string             `bson:"original_name"`
	Size         int64              `bson:"size"`
	MimeType     string             `bson:"mime_type"`
	Checksum     string             `bson:"checksum"`
	StoredOn     primitive.ObjectID `bson:"stored_on"` // Device ID where file is stored
	Status       FileStatus         `bson:"status"`
	StoragePath  string             `bson:"storage_path"` // Path on the device
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

type FileStatus string

const (
	FileStatusPending   FileStatus = "pending"
	FileStatusStored    FileStatus = "stored"
	FileStatusCorrupted FileStatus = "corrupted"
	FileStatusDeleted   FileStatus = "deleted"
)

type StoreFileRequest struct {
	Name         string `json:"name" validate:"required"`
	Size         int64  `json:"size" validate:"required,min=1"`
	MimeType     string `json:"mime_type"`
	TargetDevice string `json:"target_device,omitempty"` // Optional specific device
}

type FileMetadata struct {
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
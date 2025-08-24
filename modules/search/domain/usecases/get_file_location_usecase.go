package usecases

import (
	"context"
	"errors"

	deviceRepo "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	fileRepo "github.com/manab-pr/nebulo/modules/files/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileLocationInfo struct {
	FileID      string `json:"file_id"`
	FileName    string `json:"file_name"`
	DeviceID    string `json:"device_id"`
	DeviceName  string `json:"device_name"`
	DeviceIP    string `json:"device_ip"`
	StoragePath string `json:"storage_path"`
	Status      string `json:"status"`
}

type GetFileLocationUseCase struct {
	fileRepo   fileRepo.FileRepository
	deviceRepo deviceRepo.DeviceRepository
}

func NewGetFileLocationUseCase(fileRepo fileRepo.FileRepository, deviceRepo deviceRepo.DeviceRepository) *GetFileLocationUseCase {
	return &GetFileLocationUseCase{
		fileRepo:   fileRepo,
		deviceRepo: deviceRepo,
	}
}

func (uc *GetFileLocationUseCase) Execute(ctx context.Context, userID, fileID string) (*FileLocationInfo, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	fileObjectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return nil, errors.New("invalid file ID")
	}

	// Get file (user-scoped)
	file, err := uc.fileRepo.GetByID(ctx, userObjectID, fileObjectID)
	if err != nil {
		return nil, err
	}

	if file == nil {
		return nil, errors.New("file not found or does not belong to you")
	}

	// Get device where file is stored (user-scoped)
	device, err := uc.deviceRepo.GetByID(ctx, userObjectID, file.StoredOn)
	if err != nil {
		return nil, err
	}

	if device == nil {
		return nil, errors.New("device not found or does not belong to you")
	}

	locationInfo := &FileLocationInfo{
		FileID:      file.ID.Hex(),
		FileName:    file.OriginalName,
		DeviceID:    device.ID.Hex(),
		DeviceName:  device.Name,
		DeviceIP:    device.IPAddress,
		StoragePath: file.StoragePath,
		Status:      string(file.Status),
	}

	return locationInfo, nil
}

package usecases

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	deviceEntities "github.com/manab-pr/nebulo/modules/devices/domain/entities"
	"github.com/manab-pr/nebulo/modules/devices/domain/repository"
	"github.com/manab-pr/nebulo/modules/files/domain/entities"
	fileRepository "github.com/manab-pr/nebulo/modules/files/domain/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreFileUseCase struct {
	fileRepo   fileRepository.FileRepository
	deviceRepo repository.DeviceRepository
}

func NewStoreFileUseCase(fileRepo fileRepository.FileRepository, deviceRepo repository.DeviceRepository) *StoreFileUseCase {
	return &StoreFileUseCase{
		fileRepo:   fileRepo,
		deviceRepo: deviceRepo,
	}
}

func (uc *StoreFileUseCase) Execute(ctx context.Context, userID string, req entities.StoreFileRequest, fileData []byte) (*entities.File, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var selectedDevice *deviceEntities.Device

	// Select target device
	if req.TargetDevice != "" {
		// Use specific device
		deviceID, parseErr := primitive.ObjectIDFromHex(req.TargetDevice)
		if parseErr != nil {
			return nil, errors.New("invalid target device ID")
		}
		selectedDevice, err = uc.deviceRepo.GetByID(ctx, userObjectID, deviceID)
		if err != nil || selectedDevice == nil {
			return nil, errors.New("target device not found or does not belong to you")
		}
		if selectedDevice.Status != deviceEntities.DeviceStatusOnline {
			return nil, errors.New("target device is not online")
		}
	} else {
		// Find an online device with sufficient space
		onlineDevices, deviceErr := uc.deviceRepo.GetOnlineDevicesByUser(ctx, userObjectID)
		if deviceErr != nil {
			return nil, deviceErr
		}
		if len(onlineDevices) == 0 {
			return nil, errors.New("no online devices available")
		}

		for _, device := range onlineDevices {
			if device.AvailableStorage >= req.Size {
				selectedDevice = device
				break
			}
		}

		if selectedDevice == nil {
			return nil, errors.New("no device with sufficient storage available")
		}
	}

	// Generate unique filename and checksum
	uniqueID := uuid.New().String()
	fileName := fmt.Sprintf("%s_%s", uniqueID, req.Name)
	checksum := fmt.Sprintf("%x", sha256.Sum256(fileData))

	// Create file record
	file := &entities.File{
		ID:           primitive.NewObjectID(),
		UserID:       userObjectID,
		Name:         fileName,
		OriginalName: req.Name,
		Size:         req.Size,
		MimeType:     req.MimeType,
		Checksum:     checksum,
		StoredOn:     selectedDevice.ID,
		Status:       entities.FileStatusPending,
		StoragePath:  fmt.Sprintf("/storage/%s", fileName),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdFile, err := uc.fileRepo.Create(ctx, file)
	if err != nil {
		return nil, err
	}

	// TODO: Send file to device (this would be done via HTTP call to device's internal server)
	// For now, we'll mark it as stored immediately
	err = uc.fileRepo.UpdateStatus(ctx, userObjectID, createdFile.ID, entities.FileStatusStored)
	if err != nil {
		return nil, err
	}

	return createdFile, nil
}

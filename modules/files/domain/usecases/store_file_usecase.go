package usecases

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/manab-pr/nebulo/modules/devices/domain/repository"
	deviceEntities "github.com/manab-pr/nebulo/modules/devices/domain/entities"
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

func (uc *StoreFileUseCase) Execute(ctx context.Context, req entities.StoreFileRequest, fileData []byte) (*entities.File, error) {
	var selectedDevice *deviceEntities.Device
	var err error

	// Select target device
	if req.TargetDevice != "" {
		// Use specific device
		deviceID, err := primitive.ObjectIDFromHex(req.TargetDevice)
		if err != nil {
			return nil, errors.New("invalid target device ID")
		}
		selectedDevice, err = uc.deviceRepo.GetByID(ctx, deviceID)
		if err != nil || selectedDevice == nil {
			return nil, errors.New("target device not found")
		}
		if selectedDevice.Status != deviceEntities.DeviceStatusOnline {
			return nil, errors.New("target device is not online")
		}
	} else {
		// Find an online device with sufficient space
		onlineDevices, err := uc.deviceRepo.GetOnlineDevices(ctx)
		if err != nil {
			return nil, err
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
	checksum := fmt.Sprintf("%x", md5.Sum(fileData))

	// Create file record
	file := &entities.File{
		ID:           primitive.NewObjectID(),
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
	err = uc.fileRepo.UpdateStatus(ctx, createdFile.ID, entities.FileStatusStored)
	if err != nil {
		return nil, err
	}

	return createdFile, nil
}
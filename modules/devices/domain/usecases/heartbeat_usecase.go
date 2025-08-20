package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/devices/domain/entities"
	"github.com/manab-pr/nebulo/modules/devices/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HeartbeatUseCase struct {
	deviceRepo repository.DeviceRepository
}

func NewHeartbeatUseCase(deviceRepo repository.DeviceRepository) *HeartbeatUseCase {
	return &HeartbeatUseCase{
		deviceRepo: deviceRepo,
	}
}

func (uc *HeartbeatUseCase) Execute(ctx context.Context, req entities.DeviceHeartbeatRequest) error {
	deviceID, err := primitive.ObjectIDFromHex(req.DeviceID)
	if err != nil {
		return errors.New("invalid device ID")
	}

	// Check if device exists
	device, err := uc.deviceRepo.GetByID(ctx, deviceID)
	if err != nil || device == nil {
		return errors.New("device not found")
	}

	// Update heartbeat and storage info
	err = uc.deviceRepo.UpdateHeartbeat(ctx, deviceID, req.AvailableStorage, req.UsedStorage)
	if err != nil {
		return err
	}

	// Update device status to online
	err = uc.deviceRepo.UpdateStatus(ctx, deviceID, entities.DeviceStatusOnline)
	if err != nil {
		return err
	}

	return nil
}
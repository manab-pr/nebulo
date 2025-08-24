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

func (uc *HeartbeatUseCase) Execute(ctx context.Context, userID string, req entities.DeviceHeartbeatRequest) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	deviceID, err := primitive.ObjectIDFromHex(req.DeviceID)
	if err != nil {
		return errors.New("invalid device ID")
	}

	// Check if device exists and belongs to user
	device, err := uc.deviceRepo.GetByID(ctx, userObjectID, deviceID)
	if err != nil || device == nil {
		return errors.New("device not found or does not belong to you")
	}

	// Update heartbeat and storage info
	err = uc.deviceRepo.UpdateHeartbeat(ctx, userObjectID, deviceID, req.AvailableStorage, req.UsedStorage)
	if err != nil {
		return err
	}

	// Update device status to online
	err = uc.deviceRepo.UpdateStatus(ctx, userObjectID, deviceID, entities.DeviceStatusOnline)
	if err != nil {
		return err
	}

	return nil
}

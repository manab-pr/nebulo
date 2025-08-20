package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"
	"github.com/manab-pr/nebulo/modules/transfers/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPendingTransfersUseCase struct {
	transferRepo repository.TransferRepository
}

func NewGetPendingTransfersUseCase(transferRepo repository.TransferRepository) *GetPendingTransfersUseCase {
	return &GetPendingTransfersUseCase{
		transferRepo: transferRepo,
	}
}

func (uc *GetPendingTransfersUseCase) Execute(ctx context.Context, deviceID string) ([]*entities.Transfer, error) {
	id, err := primitive.ObjectIDFromHex(deviceID)
	if err != nil {
		return nil, errors.New("invalid device ID")
	}

	transfers, err := uc.transferRepo.GetPendingByDeviceID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update status to in_progress for retrieved transfers
	for _, transfer := range transfers {
		err = uc.transferRepo.UpdateStatus(ctx, transfer.ID, entities.TransferStatusInProgress)
		if err != nil {
			continue // Log error but continue with other transfers
		}
	}

	return transfers, nil
}
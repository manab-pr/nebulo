package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"
	"github.com/manab-pr/nebulo/modules/transfers/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CancelTransferUseCase struct {
	transferRepo repository.TransferRepository
}

func NewCancelTransferUseCase(transferRepo repository.TransferRepository) *CancelTransferUseCase {
	return &CancelTransferUseCase{
		transferRepo: transferRepo,
	}
}

func (uc *CancelTransferUseCase) Execute(ctx context.Context, transferID string) error {
	id, err := primitive.ObjectIDFromHex(transferID)
	if err != nil {
		return errors.New("invalid transfer ID")
	}

	// Check if transfer exists
	transfer, err := uc.transferRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if transfer == nil {
		return errors.New("transfer not found")
	}

	// Only allow cancellation of pending or in-progress transfers
	if transfer.Status != entities.TransferStatusPending && transfer.Status != entities.TransferStatusInProgress {
		return errors.New("cannot cancel completed or failed transfer")
	}

	// Cancel the transfer
	err = uc.transferRepo.UpdateStatus(ctx, id, entities.TransferStatusCancelled)
	if err != nil {
		return err
	}

	return nil
}
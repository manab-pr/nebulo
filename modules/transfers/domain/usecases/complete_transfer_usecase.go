package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"
	"github.com/manab-pr/nebulo/modules/transfers/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompleteTransferUseCase struct {
	transferRepo repository.TransferRepository
}

func NewCompleteTransferUseCase(transferRepo repository.TransferRepository) *CompleteTransferUseCase {
	return &CompleteTransferUseCase{
		transferRepo: transferRepo,
	}
}

func (uc *CompleteTransferUseCase) Execute(ctx context.Context, req entities.CompleteTransferRequest) error {
	transferID, err := primitive.ObjectIDFromHex(req.TransferID)
	if err != nil {
		return errors.New("invalid transfer ID")
	}

	// Check if transfer exists
	transfer, err := uc.transferRepo.GetByID(ctx, transferID)
	if err != nil {
		return err
	}

	if transfer == nil {
		return errors.New("transfer not found")
	}

	// Complete the transfer
	err = uc.transferRepo.CompleteTransfer(ctx, transferID, req.Success, req.ErrorMsg)
	if err != nil {
		return err
	}

	// If failed and under max retries, increment retry count and reset to pending
	if !req.Success && transfer.Retries < transfer.MaxRetries {
		err = uc.transferRepo.IncrementRetries(ctx, transferID)
		if err != nil {
			return err
		}
		err = uc.transferRepo.UpdateStatus(ctx, transferID, entities.TransferStatusPending)
		if err != nil {
			return err
		}
	}

	return nil
}

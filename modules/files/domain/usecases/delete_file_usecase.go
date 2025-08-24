package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/files/domain/entities"
	"github.com/manab-pr/nebulo/modules/files/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteFileUseCase struct {
	fileRepo repository.FileRepository
}

func NewDeleteFileUseCase(fileRepo repository.FileRepository) *DeleteFileUseCase {
	return &DeleteFileUseCase{
		fileRepo: fileRepo,
	}
}

func (uc *DeleteFileUseCase) Execute(ctx context.Context, userID, fileID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	fileObjectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return errors.New("invalid file ID")
	}

	// Check if file exists and belongs to user
	file, err := uc.fileRepo.GetByID(ctx, userObjectID, fileObjectID)
	if err != nil {
		return err
	}

	if file == nil {
		return errors.New("file not found or does not belong to you")
	}

	// TODO: Send delete request to the device where file is stored
	// For now, we'll just mark it as deleted and remove from database
	err = uc.fileRepo.UpdateStatus(ctx, userObjectID, fileObjectID, entities.FileStatusDeleted)
	if err != nil {
		return err
	}

	// Remove file record from database
	err = uc.fileRepo.Delete(ctx, userObjectID, fileObjectID)
	if err != nil {
		return err
	}

	return nil
}

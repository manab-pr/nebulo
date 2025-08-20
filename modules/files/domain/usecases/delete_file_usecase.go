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

func (uc *DeleteFileUseCase) Execute(ctx context.Context, fileID string) error {
	id, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return errors.New("invalid file ID")
	}

	// Check if file exists
	file, err := uc.fileRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if file == nil {
		return errors.New("file not found")
	}

	// TODO: Send delete request to the device where file is stored
	// For now, we'll just mark it as deleted and remove from database
	err = uc.fileRepo.UpdateStatus(ctx, id, entities.FileStatusDeleted)
	if err != nil {
		return err
	}

	// Remove file record from database
	err = uc.fileRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
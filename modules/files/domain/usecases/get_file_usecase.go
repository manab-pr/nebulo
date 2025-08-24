package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/nebulo/modules/files/domain/entities"
	"github.com/manab-pr/nebulo/modules/files/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFileUseCase struct {
	fileRepo repository.FileRepository
}

func NewGetFileUseCase(fileRepo repository.FileRepository) *GetFileUseCase {
	return &GetFileUseCase{
		fileRepo: fileRepo,
	}
}

func (uc *GetFileUseCase) Execute(ctx context.Context, userID, fileID string) (*entities.File, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	fileObjectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return nil, errors.New("invalid file ID")
	}

	file, err := uc.fileRepo.GetByID(ctx, userObjectID, fileObjectID)
	if err != nil {
		return nil, err
	}

	if file == nil {
		return nil, errors.New("file not found or does not belong to you")
	}

	return file, nil
}

func (uc *GetFileUseCase) GetAllFiles(ctx context.Context, userID string) ([]*entities.File, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	files, err := uc.fileRepo.GetAllByUser(ctx, userObjectID)
	if err != nil {
		return nil, err
	}

	return files, nil
}

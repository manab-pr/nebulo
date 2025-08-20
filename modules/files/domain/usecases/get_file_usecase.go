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

func (uc *GetFileUseCase) Execute(ctx context.Context, fileID string) (*entities.File, error) {
	id, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return nil, errors.New("invalid file ID")
	}

	file, err := uc.fileRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if file == nil {
		return nil, errors.New("file not found")
	}

	return file, nil
}

func (uc *GetFileUseCase) GetAllFiles(ctx context.Context) ([]*entities.File, error) {
	files, err := uc.fileRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return files, nil
}

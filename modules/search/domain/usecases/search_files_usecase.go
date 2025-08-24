package usecases

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	fileEntities "github.com/manab-pr/nebulo/modules/files/domain/entities"
	fileRepo "github.com/manab-pr/nebulo/modules/files/domain/repository"
)

type SearchFilesUseCase struct {
	fileRepo fileRepo.FileRepository
}

func NewSearchFilesUseCase(fileRepo fileRepo.FileRepository) *SearchFilesUseCase {
	return &SearchFilesUseCase{
		fileRepo: fileRepo,
	}
}

func (uc *SearchFilesUseCase) Execute(ctx context.Context, userID, query string) ([]*fileEntities.File, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	if strings.TrimSpace(query) == "" {
		// If no query provided, return all user's files
		return uc.fileRepo.GetAllByUser(ctx, userObjectID)
	}

	// Search files by name for specific user
	files, err := uc.fileRepo.SearchByNameForUser(ctx, userObjectID, query)
	if err != nil {
		return nil, err
	}

	return files, nil
}

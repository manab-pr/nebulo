package usecases

import (
	"context"
	"strings"

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

func (uc *SearchFilesUseCase) Execute(ctx context.Context, query string) ([]*fileEntities.File, error) {
	if strings.TrimSpace(query) == "" {
		// If no query provided, return all files
		return uc.fileRepo.GetAll(ctx)
	}

	// Search files by name
	files, err := uc.fileRepo.SearchByName(ctx, query)
	if err != nil {
		return nil, err
	}

	return files, nil
}

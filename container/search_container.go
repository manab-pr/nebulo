package container

import (
	deviceRepo "github.com/manab-pr/nebulo/modules/devices/domain/repository"
	fileRepo "github.com/manab-pr/nebulo/modules/files/domain/repository"
	searchUseCases "github.com/manab-pr/nebulo/modules/search/domain/usecases"
	searchHandlers "github.com/manab-pr/nebulo/modules/search/presentation/http/handlers"
)

type SearchContainer struct {
	SearchFilesUseCase    *searchUseCases.SearchFilesUseCase
	GetLocationUseCase    *searchUseCases.GetFileLocationUseCase
	Handler               *searchHandlers.SearchHandler
}

func NewSearchContainer(fileRepo fileRepo.FileRepository, deviceRepo deviceRepo.DeviceRepository) *SearchContainer {
	// Initialize use cases
	searchFilesUseCase := searchUseCases.NewSearchFilesUseCase(fileRepo)
	getLocationUseCase := searchUseCases.NewGetFileLocationUseCase(fileRepo, deviceRepo)

	// Initialize handler
	handler := searchHandlers.NewSearchHandler(
		searchFilesUseCase,
		getLocationUseCase,
	)

	return &SearchContainer{
		SearchFilesUseCase: searchFilesUseCase,
		GetLocationUseCase: getLocationUseCase,
		Handler:            handler,
	}
}
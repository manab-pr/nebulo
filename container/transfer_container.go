package container

import (
	transferRepo "github.com/manab-pr/nebulo/modules/transfers/data/mongodb/repository"
	transferUseCases "github.com/manab-pr/nebulo/modules/transfers/domain/usecases"
	transferHandlers "github.com/manab-pr/nebulo/modules/transfers/presentation/http/handlers"
	transferRepository "github.com/manab-pr/nebulo/modules/transfers/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransferContainer struct {
	Repository         transferRepository.TransferRepository
	GetPendingUseCase  *transferUseCases.GetPendingTransfersUseCase
	CompleteUseCase    *transferUseCases.CompleteTransferUseCase
	CancelUseCase      *transferUseCases.CancelTransferUseCase
	Handler            *transferHandlers.TransferHandler
}

func NewTransferContainer(db *mongo.Database) *TransferContainer {
	// Initialize repository
	repo := transferRepo.NewMongoTransferRepository(db)

	// Initialize use cases
	getPendingUseCase := transferUseCases.NewGetPendingTransfersUseCase(repo)
	completeUseCase := transferUseCases.NewCompleteTransferUseCase(repo)
	cancelUseCase := transferUseCases.NewCancelTransferUseCase(repo)

	// Initialize handler
	handler := transferHandlers.NewTransferHandler(
		getPendingUseCase,
		completeUseCase,
		cancelUseCase,
	)

	return &TransferContainer{
		Repository:        repo,
		GetPendingUseCase: getPendingUseCase,
		CompleteUseCase:   completeUseCase,
		CancelUseCase:     cancelUseCase,
		Handler:           handler,
	}
}
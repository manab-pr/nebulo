package repository

import (
	"context"

	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferRepository interface {
	Create(ctx context.Context, transfer *entities.Transfer) (*entities.Transfer, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Transfer, error)
	GetPendingByDeviceID(ctx context.Context, deviceID primitive.ObjectID) ([]*entities.Transfer, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status entities.TransferStatus) error
	CompleteTransfer(ctx context.Context, id primitive.ObjectID, success bool, errorMsg string) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]*entities.Transfer, error)
	IncrementRetries(ctx context.Context, id primitive.ObjectID) error
}

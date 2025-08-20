package repository

import (
	"context"

	"github.com/manab-pr/nebulo/modules/files/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRepository interface {
	Create(ctx context.Context, file *entities.File) (*entities.File, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*entities.File, error)
	GetAll(ctx context.Context) ([]*entities.File, error)
	GetByDeviceID(ctx context.Context, deviceID primitive.ObjectID) ([]*entities.File, error)
	Update(ctx context.Context, file *entities.File) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status entities.FileStatus) error
	SearchByName(ctx context.Context, name string) ([]*entities.File, error)
}

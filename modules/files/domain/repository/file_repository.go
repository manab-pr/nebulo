package repository

import (
	"context"

	"github.com/manab-pr/nebulo/modules/files/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRepository interface {
	Create(ctx context.Context, file *entities.File) (*entities.File, error)
	GetByID(ctx context.Context, userID, fileID primitive.ObjectID) (*entities.File, error)
	GetAllByUser(ctx context.Context, userID primitive.ObjectID) ([]*entities.File, error)
	GetByUserAndDeviceID(ctx context.Context, userID, deviceID primitive.ObjectID) ([]*entities.File, error)
	Update(ctx context.Context, file *entities.File) error
	Delete(ctx context.Context, userID, fileID primitive.ObjectID) error
	UpdateStatus(ctx context.Context, userID, fileID primitive.ObjectID, status entities.FileStatus) error
	SearchByNameForUser(ctx context.Context, userID primitive.ObjectID, name string) ([]*entities.File, error)
}

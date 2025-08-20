package repository

import (
	"context"
	"time"

	"github.com/manab-pr/nebulo/modules/transfers/data/mongodb/model"
	"github.com/manab-pr/nebulo/modules/transfers/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTransferRepository struct {
	collection *mongo.Collection
}

func NewMongoTransferRepository(db *mongo.Database) *MongoTransferRepository {
	return &MongoTransferRepository{
		collection: db.Collection("transfers"),
	}
}

func (r *MongoTransferRepository) Create(ctx context.Context, transfer *entities.Transfer) (*entities.Transfer, error) {
	transferModel := model.FromEntity(transfer)

	result, err := r.collection.InsertOne(ctx, transferModel)
	if err != nil {
		return nil, err
	}

	transfer.ID = result.InsertedID.(primitive.ObjectID)
	return transfer, nil
}

func (r *MongoTransferRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entities.Transfer, error) {
	var transferModel model.TransferModel

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transferModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return transferModel.ToEntity(), nil
}

func (r *MongoTransferRepository) GetPendingByDeviceID(ctx context.Context, deviceID primitive.ObjectID) ([]*entities.Transfer, error) {
	filter := bson.M{
		"device_id": deviceID,
		"status":    string(entities.TransferStatusPending),
	}

	opts := options.Find().SetSort(bson.D{
		{Key: "priority", Value: -1},  // Higher priority first
		{Key: "created_at", Value: 1}, // Older first within same priority
	})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transfers []*entities.Transfer
	for cursor.Next(ctx) {
		var transferModel model.TransferModel
		if err := cursor.Decode(&transferModel); err != nil {
			continue
		}
		transfers = append(transfers, transferModel.ToEntity())
	}

	return transfers, nil
}

func (r *MongoTransferRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status entities.TransferStatus) error {
	update := bson.M{
		"$set": bson.M{
			"status":     string(status),
			"updated_at": time.Now(),
		},
	}

	if status == entities.TransferStatusInProgress {
		now := time.Now()
		update["$set"].(bson.M)["started_at"] = &now
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func (r *MongoTransferRepository) CompleteTransfer(ctx context.Context, id primitive.ObjectID, success bool, errorMsg string) error {
	status := entities.TransferStatusCompleted
	if !success {
		status = entities.TransferStatusFailed
	}

	now := time.Now()
	update := bson.M{
		"$set": bson.M{
			"status":       string(status),
			"updated_at":   now,
			"completed_at": &now,
			"error_msg":    errorMsg,
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func (r *MongoTransferRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoTransferRepository) GetAll(ctx context.Context) ([]*entities.Transfer, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transfers []*entities.Transfer
	for cursor.Next(ctx) {
		var transferModel model.TransferModel
		if err := cursor.Decode(&transferModel); err != nil {
			continue
		}
		transfers = append(transfers, transferModel.ToEntity())
	}

	return transfers, nil
}

func (r *MongoTransferRepository) IncrementRetries(ctx context.Context, id primitive.ObjectID) error {
	update := bson.M{
		"$inc": bson.M{"retries": 1},
		"$set": bson.M{"updated_at": time.Now()},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

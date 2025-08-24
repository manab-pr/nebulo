package repository

import (
	"context"
	"time"

	"github.com/manab-pr/nebulo/modules/devices/data/mongodb/model"
	"github.com/manab-pr/nebulo/modules/devices/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDeviceRepository struct {
	collection *mongo.Collection
}

func NewMongoDeviceRepository(db *mongo.Database) *MongoDeviceRepository {
	return &MongoDeviceRepository{
		collection: db.Collection("devices"),
	}
}

func (r *MongoDeviceRepository) Create(ctx context.Context, device *entities.Device) (*entities.Device, error) {
	deviceModel := model.FromEntity(device)

	result, err := r.collection.InsertOne(ctx, deviceModel)
	if err != nil {
		return nil, err
	}

	device.ID = result.InsertedID.(primitive.ObjectID)
	return device, nil
}

func (r *MongoDeviceRepository) GetByID(ctx context.Context, userID, deviceID primitive.ObjectID) (*entities.Device, error) {
	var deviceModel model.DeviceModel

	err := r.collection.FindOne(ctx, bson.M{"_id": deviceID, "user_id": userID}).Decode(&deviceModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return deviceModel.ToEntity(), nil
}

func (r *MongoDeviceRepository) GetByIPAddress(ctx context.Context, userID primitive.ObjectID, ipAddress string) (*entities.Device, error) {
	var deviceModel model.DeviceModel

	err := r.collection.FindOne(ctx, bson.M{"ip_address": ipAddress, "user_id": userID}).Decode(&deviceModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return deviceModel.ToEntity(), nil
}

func (r *MongoDeviceRepository) GetAllByUser(ctx context.Context, userID primitive.ObjectID) ([]*entities.Device, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var devices []*entities.Device
	for cursor.Next(ctx) {
		var deviceModel model.DeviceModel
		if err := cursor.Decode(&deviceModel); err != nil {
			continue
		}
		devices = append(devices, deviceModel.ToEntity())
	}

	return devices, nil
}

func (r *MongoDeviceRepository) GetOnlineDevicesByUser(ctx context.Context, userID primitive.ObjectID) ([]*entities.Device, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"status": string(entities.DeviceStatusOnline), "user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var devices []*entities.Device
	for cursor.Next(ctx) {
		var deviceModel model.DeviceModel
		if err := cursor.Decode(&deviceModel); err != nil {
			continue
		}
		devices = append(devices, deviceModel.ToEntity())
	}

	return devices, nil
}

func (r *MongoDeviceRepository) Update(ctx context.Context, device *entities.Device) error {
	deviceModel := model.FromEntity(device)
	deviceModel.UpdatedAt = time.Now()

	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": device.ID, "user_id": device.UserID}, deviceModel)
	return err
}

func (r *MongoDeviceRepository) UpdateHeartbeat(
	ctx context.Context, userID, deviceID primitive.ObjectID, availableStorage, usedStorage int64,
) error {
	update := bson.M{
		"$set": bson.M{
			"available_storage": availableStorage,
			"used_storage":      usedStorage,
			"last_heartbeat":    time.Now(),
			"updated_at":        time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": deviceID, "user_id": userID}, update)
	return err
}

func (r *MongoDeviceRepository) Delete(ctx context.Context, userID, deviceID primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": deviceID, "user_id": userID})
	return err
}

func (r *MongoDeviceRepository) UpdateStatus(ctx context.Context, userID, deviceID primitive.ObjectID, status entities.DeviceStatus) error {
	update := bson.M{
		"$set": bson.M{
			"status":     string(status),
			"updated_at": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": deviceID, "user_id": userID}, update)
	return err
}

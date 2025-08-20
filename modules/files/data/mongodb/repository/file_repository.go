package repository

import (
	"context"
	"time"

	"github.com/manab-pr/nebulo/modules/files/data/mongodb/model"
	"github.com/manab-pr/nebulo/modules/files/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFileRepository struct {
	collection *mongo.Collection
}

func NewMongoFileRepository(db *mongo.Database) *MongoFileRepository {
	return &MongoFileRepository{
		collection: db.Collection("files"),
	}
}

func (r *MongoFileRepository) Create(ctx context.Context, file *entities.File) (*entities.File, error) {
	fileModel := model.FromEntity(file)

	result, err := r.collection.InsertOne(ctx, fileModel)
	if err != nil {
		return nil, err
	}

	file.ID = result.InsertedID.(primitive.ObjectID)
	return file, nil
}

func (r *MongoFileRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entities.File, error) {
	var fileModel model.FileModel

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&fileModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return fileModel.ToEntity(), nil
}

func (r *MongoFileRepository) GetAll(ctx context.Context) ([]*entities.File, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []*entities.File
	for cursor.Next(ctx) {
		var fileModel model.FileModel
		if err := cursor.Decode(&fileModel); err != nil {
			continue
		}
		files = append(files, fileModel.ToEntity())
	}

	return files, nil
}

func (r *MongoFileRepository) GetByDeviceID(ctx context.Context, deviceID primitive.ObjectID) ([]*entities.File, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"stored_on": deviceID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []*entities.File
	for cursor.Next(ctx) {
		var fileModel model.FileModel
		if err := cursor.Decode(&fileModel); err != nil {
			continue
		}
		files = append(files, fileModel.ToEntity())
	}

	return files, nil
}

func (r *MongoFileRepository) Update(ctx context.Context, file *entities.File) error {
	fileModel := model.FromEntity(file)
	fileModel.UpdatedAt = time.Now()

	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": file.ID}, fileModel)
	return err
}

func (r *MongoFileRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoFileRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status entities.FileStatus) error {
	update := bson.M{
		"$set": bson.M{
			"status":     string(status),
			"updated_at": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func (r *MongoFileRepository) SearchByName(ctx context.Context, name string) ([]*entities.File, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": name, "$options": "i"}},
			{"original_name": bson.M{"$regex": name, "$options": "i"}},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []*entities.File
	for cursor.Next(ctx) {
		var fileModel model.FileModel
		if err := cursor.Decode(&fileModel); err != nil {
			continue
		}
		files = append(files, fileModel.ToEntity())
	}

	return files, nil
}

package indexes

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDeviceIndexes(collection *mongo.Collection) error {
	ctx := context.Background()

	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "ip_address", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "status", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "last_heartbeat", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("Failed to create device indexes: %v", err)
		return err
	}

	log.Println("Device indexes created successfully")
	return nil
}
package indexes

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/manab-pr/nebulo/modules/users/domain/constants"
)

func CreateUserIndexes(db *mongo.Database) error {
	collection := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), constants.IndexTimeoutSeconds*time.Second)
	defer cancel()

	// Create unique index on phone_number
	phoneIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "phone_number", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// Create compound index on phone_number and otp for faster OTP lookups
	otpIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "phone_number", Value: 1},
			{Key: "otp", Value: 1},
		},
	}

	// Create TTL index on otp_expiry
	ttlIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "otp_expiry", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0), // Expire based on the date value
	}

	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		phoneIndex,
		otpIndex,
		ttlIndex,
	})

	return err
}

package repository

import (
	"context"
	"time"

	"github.com/manab-pr/nebulo/modules/users/data/mongodb/model"
	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entities.User) error {
	userModel := model.FromEntity(user)
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, userModel)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt
	return nil
}

func (r *userRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.User, error) {
	var userModel model.UserModel
	err := r.collection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&userModel)
	if err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID string) (*entities.User, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	var userModel model.UserModel
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&userModel)
	if err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	user.UpdatedAt = time.Now()
	userModel := model.FromEntity(user)

	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": user.ID}, userModel)
	return err
}

func (r *userRepository) UpdateOTP(ctx context.Context, phoneNumber, otp string, expiry time.Time) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"phone_number": phoneNumber},
		bson.M{
			"$set": bson.M{
				"otp":        otp,
				"otp_expiry": expiry,
				"updated_at": time.Now(),
			},
		},
	)
	return err
}

func (r *userRepository) ClearOTP(ctx context.Context, phoneNumber string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"phone_number": phoneNumber},
		bson.M{
			"$unset": bson.M{
				"otp":        "",
				"otp_expiry": "",
			},
			"$set": bson.M{
				"is_verified": true,
				"updated_at":  time.Now(),
			},
		},
	)
	return err
}
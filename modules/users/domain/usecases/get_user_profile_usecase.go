package usecases

import (
	"context"

	"github.com/manab-pr/nebulo/modules/users/domain/entities"
	"github.com/manab-pr/nebulo/modules/users/domain/repository"
)

type GetUserProfileUseCase struct {
	userRepo repository.UserRepository
}

func NewGetUserProfileUseCase(userRepo repository.UserRepository) *GetUserProfileUseCase {
	return &GetUserProfileUseCase{
		userRepo: userRepo,
	}
}

func (uc *GetUserProfileUseCase) Execute(ctx context.Context, userID string) (*entities.UserProfile, error) {
	user, err := uc.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entities.UserProfile{
		ID:          user.ID.Hex(),
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		IsVerified:  user.IsVerified,
		CreatedAt:   user.CreatedAt,
	}, nil
}

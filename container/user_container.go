package container

import (
	userRepo "github.com/manab-pr/nebulo/modules/users/data/mongodb/repository"
	"github.com/manab-pr/nebulo/modules/users/domain/usecases"
	"github.com/manab-pr/nebulo/modules/users/presentation/http/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserContainer struct {
	RegisterUseCase       *usecases.RegisterUserUseCase
	LoginUseCase          *usecases.LoginUserUseCase
	VerifyOTPUseCase      *usecases.VerifyOTPUseCase
	GetUserProfileUseCase *usecases.GetUserProfileUseCase
	UserHandler           *handlers.UserHandler
}

func NewUserContainer(db *mongo.Database) *UserContainer {
	// Repository
	userRepository := userRepo.NewUserRepository(db)

	// Use cases
	registerUseCase := usecases.NewRegisterUserUseCase(userRepository)
	loginUseCase := usecases.NewLoginUserUseCase(userRepository)
	verifyOTPUseCase := usecases.NewVerifyOTPUseCase(userRepository)
	getUserProfileUseCase := usecases.NewGetUserProfileUseCase(userRepository)

	// Handler
	userHandler := handlers.NewUserHandler(
		registerUseCase,
		loginUseCase,
		verifyOTPUseCase,
		getUserProfileUseCase,
	)

	return &UserContainer{
		RegisterUseCase:       registerUseCase,
		LoginUseCase:          loginUseCase,
		VerifyOTPUseCase:      verifyOTPUseCase,
		GetUserProfileUseCase: getUserProfileUseCase,
		UserHandler:           userHandler,
	}
}
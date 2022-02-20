package injector

import (
	"GoBlogClean/auth/repository"
	"GoBlogClean/auth/handler"
	"GoBlogClean/auth/usecase"
	"GoBlogClean/models"
)

func InjectUserRepository() models.UserRepository {
	return repository.NewUserRepository(InjectDB())
}

func InjectUserUsecase() models.UserUsecase {
	return usecase.NewUserUsecase(InjectUserRepository())
}

func InjectUserHandler() handler.UserHandler {
	return handler.NewUserHandler(InjectUserUsecase())
}
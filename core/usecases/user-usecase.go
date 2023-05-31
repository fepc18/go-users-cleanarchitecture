package usecases

import (
	"errors"

	"github.com/fepc18/go-users-cleanarchitecture/core/domain/models"
)

type UserUseCase struct {
	UserRepo models.IUserRepository
	//User *models.User
}

// constructor
func NewUserUseCase(ur models.IUserRepository) UserUseCase {
	return UserUseCase{UserRepo: ur}
}

// methods
func (uc *UserUseCase) Create(user *models.User) (*models.User, error) {
	userCreated, status, err := uc.UserRepo.Create(user)

	//_, status, err := mongodb.InsertRegister(t)
	if err != nil {
		//http.Error(w, "An error occurred while trying to register the user "+err.Error(), http.StatusInternalServerError)
		return user, err
	}
	if status == false {
		return user, errors.New("Sample Error")
	}

	return userCreated, nil
}

/*
func (uc *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	createdUser, err := uc.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}*/

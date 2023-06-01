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
	if err != nil {
		return user, err
	}
	if status == false {
		return user, errors.New("An error occurred while trying to register the user")
	}

	return userCreated, nil
}

// Use case  CheckUserExist is the function that allows to know if the user already exists in the database
func (uc *UserUseCase) CheckUserExist(email string) (bool, error) {
	_, encontrado, err := uc.UserRepo.CheckUserExist(email)
	if err != nil {
		return false, err
	}
	return encontrado, nil
}

package usecases

import (
	"errors"

	"github.com/fepc18/go-users-cleanarchitecture/core/domain/models"
	"github.com/fepc18/go-users-cleanarchitecture/infraestructure/persistence/mongodb"
)

type UserUseCase struct {
	value string
	//UserRepo domain.UserRepository
	//User *models.User
}

// constructor
// func NewUserUseCase(ur domain.UserRepository) UserUseCase {
func NewUserUseCase() UserUseCase {
	//func NewUserUseCase(User *models.User) (*models.User, error) *UserUseCase {
	//return User, nil
	return UserUseCase{}
}

// methods
func (uc *UserUseCase) Create(t models.User) (string, error) {

	_, status, err := mongodb.InsertRegister(t)
	if err != nil {
		//http.Error(w, "An error occurred while trying to register the user "+err.Error(), http.StatusInternalServerError)
		return "", err
	}
	if status == false {
		return "", errors.New("Sample Error")
	}
	return "User created ", nil
}

/*
func (uc *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	createdUser, err := uc.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}*/

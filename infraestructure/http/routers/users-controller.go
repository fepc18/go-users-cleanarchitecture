package routers

import (
	"encoding/json"
	"net/http"

	//"github.com/fepc18/twiter/bd"
	"github.com/fepc18/go-users-cleanarchitecture/core/domain/models"
	"github.com/fepc18/go-users-cleanarchitecture/core/usecases"
	"github.com/fepc18/go-users-cleanarchitecture/infraestructure/http/responses"
	"github.com/gorilla/mux"
)

type UserController struct {
	UserUseCase *usecases.UserUseCase //user Service
}

// factory
func NewUserController(r *mux.Router, uc usecases.UserUseCase) {
	controller := UserController{UserUseCase: &uc}
	r.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), http.StatusBadRequest)
		return
	}

	validated, errors := ValidateUser(t)
	if !validated {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.NewErrorResponse("1234", http.StatusBadRequest, "Invalid request parameters.", errors))
		return
	}
	encontrado, _ := uc.UserUseCase.CheckUserExist(t.Email)
	if encontrado == true {
		//http.StatusConflict(w)
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(responses.NewErrorResponse("1234", http.StatusConflict, "User already exists", nil))
		return
	}

	data, err := uc.UserUseCase.Create(&t)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.NewErrorResponse("1234", http.StatusInternalServerError, "An error occurred while trying to register the user "+err.Error(), nil))
		return
	}

	/*
		_, status, err := bd.InsertRegister(t)
		if status == false {
			http.Error(w, "An error occurred while trying to register the user ", http.StatusInternalServerError)
			return
		}
	*/
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.NewSuccessResponse(http.StatusCreated, data))

}

func ValidateUser(t models.User) (bool, []responses.Error) {
	var errors []responses.Error

	if len(t.Email) == 0 {
		errors = append(errors, responses.Error{ErrorCode: responses.MissingField, Message: "Email is required"})

	}
	if len(t.Password) < 6 {
		errors = append(errors, responses.Error{ErrorCode: responses.InvalidField, Message: "Password must be at least 6 characters"})
	}
	if len(errors) > 0 {
		return false, errors
	}
	return true, nil
}

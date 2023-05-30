package routers

import (
	"encoding/json"
	"net/http"

	//"github.com/fepc18/twiter/bd"
	"github.com/fepc18/go-users-cleanarchitecture/core/domain/models"
	"github.com/fepc18/go-users-cleanarchitecture/core/usecases"
	"github.com/gorilla/mux"
)

type UserController struct {
	UserUseCase *usecases.UserUseCase //user Service
}

// constructor
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

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	resu, _ := uc.UserUseCase.Create(t)
	json.NewEncoder(w).Encode(resu)
	/*_, encontrado, _ := bd.CheckUserExist(t.Email)
	if encontrado == true {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to register the user "+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(w, "An error occurred while trying to register the user ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)*/

}

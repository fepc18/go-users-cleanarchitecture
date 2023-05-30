package handlers

import (
	"log"
	"net/http"
	"os"

	//"github.com/fepc18/go-users-cleanarchitecture/middleware"

	"github.com/fepc18/go-users-cleanarchitecture/core/usecases"
	"github.com/fepc18/go-users-cleanarchitecture/infraestructure/http/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {

	//userRepo := repository.NewUserRepo(conf.DBUrl)

	// Use Cases
	userUseCase := usecases.NewUserUseCase()

	// Controllers
	//r := http.NewServeMux()
	router := mux.NewRouter()
	routers.NewUserController(router, userUseCase)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

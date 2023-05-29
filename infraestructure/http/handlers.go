package handlers

import (
	"log"
	"net/http"
	"os"

	//"github.com/fepc18/go-users-cleanarchitecture/middleware"
	"github.com/fepc18/go-users-cleanarchitecture/infraestructure/http/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	//router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/register", routers.Register).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

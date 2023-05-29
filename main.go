package main

import (
	//"github.com/fepc18/twiter/bd"
	handlers "github.com/fepc18/go-users-cleanarchitecture/infraestructure/http"
)

func main() {
	/*if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}*/
	handlers.Handlers()
}

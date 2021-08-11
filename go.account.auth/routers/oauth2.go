package routers

import (
	"github.com/gorilla/mux"
	"go.account.auth/controllers"
)

func AddOauth2Routers(router *mux.Router) *mux.Router {

	router.HandleFunc("/login", controllers.Login).Methods("GET")

	return router
}

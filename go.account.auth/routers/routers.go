package routers

import "github.com/gorilla/mux"

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// set url prefix
	prefix := "/api/v1/auth"
	s := r.PathPrefix(prefix).Subrouter()

	// add handlers
	s = AddOauth2Routers(s)

	return s
}

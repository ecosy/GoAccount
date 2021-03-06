package main

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/errors"
	"github.com/go-oauth2/oauth2/manage"
	"github.com/go-oauth2/oauth2/models"
	"github.com/go-oauth2/oauth2/server"
	"github.com/go-oauth2/oauth2/store"
)

func main() {
	manager := manage.NewDefaultManager()

	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("41c34055-0f30-479c-9319-9f06a36f4b11", &models.Client{
		ID:     "41c34055-0f30-479c-9319-9f06a36f4b11",
		Secret: "NDFjMzQwNTUtMGYzMC00NzljLTkzMTktOWYwNmEzNmY0YjEx",
		Domain: "http://127.0.0.1",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error: ", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error: ", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}

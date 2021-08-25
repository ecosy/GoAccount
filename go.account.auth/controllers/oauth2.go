package controllers

import (
	"log"
	"net/http"

	"github.com/unrolled/render"
)

var rd *render.Render

func init() {
	rd = render.New(render.Options{
		Directory:  "template",
		Extensions: []string{".html"},
	})
}

// GET http://127.0.0.1:9000/api/v1/auth/login
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("/Login start....")

	rd.HTML(w, http.StatusOK, "login", "")
}

// GET http://127.0.0.1:9000/api/v1/auth/authorize
func Authorize(w http.ResponseWriter, r *http.Request) {
	// parse parameters
	authorization, err := ParseAuthRequest(r)

	// validate parameters
	// client_id, redirect_uri, response_type, scope 유효성 검사

}

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

// http://127.0.0.1:9000/api/v1/auth/login
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("/Login start....")

	rd.HTML(w, http.StatusOK, "login", "")
}

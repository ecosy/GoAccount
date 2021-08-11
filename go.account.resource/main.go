package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler start....")
	w.Write([]byte("home hanlder\n"))
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ClientHandler start....")

	rd.HTML(w, http.StatusOK, "start", "")
}

func main() {
	log.Println("main start...")

	rd = render.New(render.Options{
		Directory:  "template",
		Extensions: []string{".html"},
	})

	r := mux.NewRouter()
	r.HandleFunc("/abc", HomeHandler)
	r.HandleFunc("/client/example", ClientHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

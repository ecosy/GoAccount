package main

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler start....")
	w.Write([]byte("home hanlder\n"))
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ClientHandler start....")

	p := path.Dir("./pages/client/start.html")
	w.Header().Set("Content-type", "text/html")

	log.Println(p)
	log.Println(w)

	http.ServeFile(w, r, p)
}

func main() {
	log.Println("main start...")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/client", ClientHandler).Methods("GET")

	srv := &http.Server{
		Addr: "0.0.0.0:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Fatal(srv.ListenAndServe())
}

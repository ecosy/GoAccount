package main

import (
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
	"go.account.auth/routers"
)

func main() {
	log.Println("main start...")

	r := routers.InitRouter()

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	srv := &http.Server{
		Handler:      n,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

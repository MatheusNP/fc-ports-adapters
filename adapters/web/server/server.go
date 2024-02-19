package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MatheusNP/fc-ports-adapters/adapters/web/handler"
	"github.com/MatheusNP/fc-ports-adapters/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

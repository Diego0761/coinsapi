package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/diego0761/coinsapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting API")

	err := http.ListenAndServe("localhost:8000", r)
	
	if err != nil {
		log.Error(err)
	}
}
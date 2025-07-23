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

	router := chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting API")

	err := http.ListenAndServe("localhost:8000", router)
	
	if err != nil {
		log.Error(err)
	}
}
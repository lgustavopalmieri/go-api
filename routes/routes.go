package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lgustavopalmieri/go-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":5000", r))
}

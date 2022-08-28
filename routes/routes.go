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
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	log.Fatal(http.ListenAndServe(":5000", r))
}

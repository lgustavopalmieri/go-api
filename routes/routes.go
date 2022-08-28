package routes

import (
	"log"
	"net/http"

	"github.com/lgustavopalmieri/go-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

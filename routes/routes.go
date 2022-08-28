package routes

import (
	"log"
	"net/http"

	"github.com/lgustavopalmieri/go-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

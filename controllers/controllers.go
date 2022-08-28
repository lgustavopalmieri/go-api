package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lgustavopalmieri/go-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		//strconv.Itoa converts Int for string
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}

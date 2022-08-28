package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lgustavopalmieri/go-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

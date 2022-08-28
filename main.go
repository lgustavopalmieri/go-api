package main

import (
	"fmt"

	"github.com/lgustavopalmieri/go-api/models"
	"github.com/lgustavopalmieri/go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "First name", History: "Chapter One"},
		{Name: "Second name", History: "Chapter Two"},
	}

	fmt.Println("FirstSteps")
	routes.HandleRequest()
}

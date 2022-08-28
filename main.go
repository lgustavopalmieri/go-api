package main

import (
	"fmt"

	"github.com/lgustavopalmieri/go-api/models"
	"github.com/lgustavopalmieri/go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "First name", History: "Chapter One"},
		{Id: 2, Name: "Second name", History: "Chapter Two"},
	}

	fmt.Println("FirstSteps")
	routes.HandleRequest()
}

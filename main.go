package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	fmt.Println("FirstSteps")
	HandleRequest()
}

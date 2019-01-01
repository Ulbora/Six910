package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// just the start of Ulbora Cart Server
	// This is where REST services will be defined.
	fmt.Println("Ulbora Cart is running on port 3000!")
	router := mux.NewRouter()
	http.ListenAndServe(":3000", router)
}

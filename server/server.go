package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// just the start of Ulbora Cart
	fmt.Println("Ulbora Cart is running on port 3000!")
	router := mux.NewRouter()
	http.ListenAndServe(":3000", router)
}

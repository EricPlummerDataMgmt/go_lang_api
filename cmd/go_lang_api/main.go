package main

import (
	"fmt"
	"net/http"
	"github.com/ericplummerdatamgmt/go_lang_api/internal/routes"
)

func main() {
	fmt.Println("Hello, World!")

	router := routes.NewRouter()

	port := 8080
	addr :- fmt.Sprintf(":%d", port)

	fmt.Printf("Server Listening on http://localhost%s...\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

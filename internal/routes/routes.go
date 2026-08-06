package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHander)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln((w, "Welcome to the Home Page"))
	

package main

import (
	"net/http"
)

func handleErr(w http.ResponseWriter, r *http.Request){
	respondWithJson(w, 400, "Some thing went wrong")
}
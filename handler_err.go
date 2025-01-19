package main

import (
	"net/http"
)

func handleErr(w http.ResponseWriter, r *http.Request){
	responWithJson(w, 400, "Some thing went wrong")
}
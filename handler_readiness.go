package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	responWithJson(w,200,struct{}{})
}
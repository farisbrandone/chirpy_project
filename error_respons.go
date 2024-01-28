package main

import (
	"net/http"
	"encoding/json"
	"log"
)
type ParamsErrors struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	
	myError:=ParamsErrors{
		Error:msg,
	}
	dat, err := json.Marshal(myError)
	if err !=nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(code)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)  
}
package main

import (
	"net/http"
	"encoding/json"
	"github.com/farisbrandone/chirpy_project/internal/database"
)

/* type myUser database.User
type myChirps database.Chirp */

/*  func NewChirp(Body string, Id int) *Chirp {
	return &Chirp{
		Body: Body,
		Id:   Id,
	}
}

func NewUser(Email string, Id int) *User {
	return &User{
		Email: Email,
		Id:    Id,
	}
}  */

type UserDisplay struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	IsChirpyRed bool `json:"is_chirpy_red"`
} 
type UserRefresh struct {
	Token string `json:"token"`
} 

type UserRevoke struct {
	RevokeToken string `json:"refresh_token"`
} 

func  respondWithJSON(w http.ResponseWriter, code int, payload database.Chirp){
		dat, err := json.Marshal(payload)
		if err !=nil {
				
				w.WriteHeader(code)
				return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(dat)  
		
	}

func  respondWithJSONUser(w http.ResponseWriter, code int, payload database.User){
	
	    userDisplay:=UserDisplay{
			Email:payload.Email,
			Id:payload.Id,
			IsChirpyRed:payload.IsChirpyRed,
		}
		dat, err := json.Marshal( userDisplay)
		if err !=nil {
				
				w.WriteHeader(code)
				return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(dat)  
		
	}

func  respondWithJSONUserLogin(w http.ResponseWriter, code int, payload ResToken){
	
	    
		dat, err := json.Marshal(payload)
		if err !=nil {
				
				w.WriteHeader(code)
				return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(dat)  
		
	}

	func  respondWithJSONUserUpdate(w http.ResponseWriter, code int, payload database.User){
	
	    myResponse:=UserDisplay{
			Email:payload.Email,
			Id:payload.Id,
		}
		dat, err := json.Marshal(myResponse)
		if err !=nil {
				
				w.WriteHeader(code)
				return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(dat)  
		
	}

	func  respondWithJSONUserRefresh(w http.ResponseWriter, code int, payload string){
	
	    myResponse1:=UserRefresh{
			Token:payload,
		}
		
		dat1, err12 := json.Marshal(myResponse1)
		if err12 !=nil {
				
				w.WriteHeader(500)
				return
			}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(dat1)  
		
	}

	func  respondWithJSONUserRevoke(w http.ResponseWriter, code int, payload string){
	
	    myResponse1:= UserRevoke{
			RevokeToken	:payload,
		}
		
		dat1, err12 := json.Marshal(myResponse1)
		if err12 !=nil {
				
				w.WriteHeader(500)
				return
			}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(dat1)  
		
	}


	

	
		
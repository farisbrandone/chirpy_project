package main

import (
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"strconv"

)




func (cfg *dbConfig) handlerGetIdUserForHeaders(r *http.Request)(int, string){

    headers:=r.Header.Get("Authorization")
		if headers == "" {
			
			return 0,"deuxième erreur"
		}
		tokenString:=strings.Split(headers, " ")
		if len(tokenString) < 2 || tokenString[0] != "Bearer" {
			
			return 0, "troisieme erreur"
		}
       
		claimsStruct := jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString[1], &claimsStruct, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.config.jwtSecret), nil})
		if err != nil {
			return 0, "Quatrième erreur"
		} 
		Id1, err := token.Claims.GetSubject() 
		if err != nil {
			return 0, "cinquième erreur"
		}  
			 id, err := strconv.Atoi(Id1)
			 if err != nil {
				return 0, "sixième erreurs"
			 }

			 return id, ""


}


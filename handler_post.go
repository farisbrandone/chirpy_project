package main

import(
	"net/http"
	"encoding/json"
	"log"
	"strconv"
	"github.com/go-chi/chi/v5"
	"strings"
	//"fmt"
	"github.com/farisbrandone/chirpy_project/internal/database"
	"github.com/golang-jwt/jwt/v5"
	
)




func ( cfg dbConfig ) handlerPostVal(w http.ResponseWriter, r *http.Request){
   
	decoder := json.NewDecoder(r.Body)
    param := Bol{}
    err := decoder.Decode(&param)
	
    if err != nil {
		respondWithError(w, 500,"Something went wrong")
		return
    }
	id,stringErr:=cfg.handlerGetIdUserForHeaders(r)
	if id==0 {
		respondWithError(w, 500,stringErr)
		return
	}

	result,err:=cfg.db.CreateChirp(param.Body, id)
	if err != nil {
		respondWithError(w, 500,"Something went wrong")
		return
    }
	
	respondWithJSON(w, 201, result)
	}

	func (cfg *dbConfig) handlerPostUser(w http.ResponseWriter, r *http.Request){
		
		decoder := json.NewDecoder(r.Body)
		param := User{}
		err := decoder.Decode(&param)
		//fmt.Println(params)
		if err != nil {
			respondWithError(w, 500,"Something went wrong")
			return
		}
		bcriptPassword,err:=GenereCriptPassword(param.Password)
		if err!=nil {
			
			w.WriteHeader(500)
			return
		}
		paramParameter:=database.UserEmailPassword{
            Email:param.Email,
			Password:bcriptPassword,
		}
		result,err:=cfg.db.CreateUser(paramParameter)
	if err != nil {
		respondWithError(w, 500,"Something went wrong")
		return
    }
	respondWithJSONUser(w, 201, result)
		}


	func (cfg *dbConfig) handlerPostValGet(w http.ResponseWriter, r *http.Request){
		results,err:=cfg.db.GetChirps()
			if err!=nil {
				respondWithError(w, 500,"Something went wrong when we create data")
				return
			 }
		a:=len(results)
		log.Println("myvalue is :", a)
		if a==0 {
			dat, err := json.Marshal(results)
			if err !=nil {
					w.WriteHeader(500)
					return
			}
		
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(dat)  
            return
		}

		myResult:=make([]database.Chirp,0)
		
		s := r.URL.Query().Get("author_id")
		t:=  r.URL.Query().Get("sort")
		log.Println("my value s is ohh:", s+"k")
		if s !="" {
			id, err := strconv.Atoi(s)
			if err != nil {
				respondWithError(w, 500,"Something went wrong or your id is not a string")
				return
			}
			
			 if t=="desc" {
				for i:=a-1; i>0; i-- {
					if results[i].AuthorId == id {
						myResult=append(myResult,results[i])
					}
				}
			 }
			 if t=="acs" || t=="" {
				for i:=0; i<a; i++ {
					if results[i].AuthorId == id {
						myResult=append(myResult,results[i])
					}
				}
			 }
			
			if len(myResult)==0 {
				respondWithError(w, 500,"This author ID doesn't exist")
				return
			}
			dat, err := json.Marshal(myResult)
			if err !=nil {
					
					w.WriteHeader(500)
					return
			}
		
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(dat)  
            return
		}

		if t=="desc" {
			for i:=a-1; i>=0; i-- {
				myResult=append(myResult,results[i])
			}
		 }
		 if t=="asc" || t=="" {
			myResult=append(myResult, results...)
				
			}
		 
		
		dat, err := json.Marshal(myResult)
			if err !=nil {
					
					w.WriteHeader(500)
					return
			}
		
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(dat)  

		}


		func (/* params *ArrayBol */cfg *dbConfig) handlerPostValGetId(w http.ResponseWriter, r *http.Request){
			/* myDb, err:=NewDB("file.json")
			if err!=nil {
				respondWithError(w, 500,"Something went wrong when we create data")
				return
			 }
			idString:=chi.URLParam(r, "chirpId") 
			  fmt.Println("dododododododo", idString)
			  id, err := strconv.Atoi(idString)
			  index:=-1
			  if err != nil {
				log.Printf("Error marshalling JSON: %s", err)
				w.WriteHeader(500)
				return
			}
			if len(*params)==0 {
				respondWithError(w, 400,"no data inside database so enter data with post request")
				return
			}
			for i:=0; i<len(*params);i++ {
				if (*params)[i].Id==id {
					fmt.Println("koukou",(*params)[i].Id)
                     index=i
				}
			}
			if index==-1{
                respondWithError(w, 404,"no value for this parameter")
				return
			} */

			idString:=chi.URLParam(r, "chirpId") 
			id, err := strconv.Atoi(idString)
			if err != nil {
			  
			  w.WriteHeader(500)
			  return
		  }


			result,err:=cfg.db.GetChirp(id)
			if err!=nil {
			   respondWithError(w, 500,"Something went wrong ")
			   return
			}
			dat, err := json.Marshal(result) //(*params)[index]
				if err !=nil {
						
						w.WriteHeader(500)
						return
				}
			
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(dat)  
	
			}


	func (cfg *dbConfig)handlerLoginUser(w http.ResponseWriter, r *http.Request){
		decoder := json.NewDecoder(r.Body)
        param := database.UserEmailPassword{}
         err := decoder.Decode(&param)
		 if err != nil {
			respondWithError(w, 500,"Something went wrong")
			return
		}
		
		PasswordEmailCompare, err1:=cfg.db.LoginUser(param)
		
		if err1!=nil {
			
			if err.Error()=="password is not correct"{
				respondWithError(w, 401,"password is not correct")
				return
			}
		
			respondWithError(w, 401,"this email doesn't exist")
			return
		}
		myResponseToken,err3:=cfg.GetToken(param, PasswordEmailCompare)
		if err3!=nil {
			respondWithError(w, 500,"Something went wrong")
			return
		}

		respondWithJSONUserLogin(w, 200, myResponseToken)

	}


	func(cfg *dbConfig) handlerPutUser(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
        param := database.UserEmailPassword{}
         err := decoder.Decode(&param)
		 if err != nil {
			respondWithError(w, 500,"premiere erreur")
			return
		}
		
		headers:=r.Header.Get("Authorization")
		if headers == "" {
			respondWithError(w, 500,"deuxième erreur")
			return
		}
		tokenString:=strings.Split(headers, " ")
		if len(tokenString) < 2 || tokenString[0] != "Bearer" {
			respondWithError(w, 500,"troisieme erreur")
			return
		}
       
		claimsStruct := jwt.RegisteredClaims{}
	
		token, err := jwt.ParseWithClaims(tokenString[1], &claimsStruct, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.config.jwtSecret), nil})
		if err != nil {
			respondWithError(w, 401,"Quatrième erreur")
			return
		} 
		Id1, err := token.Claims.GetSubject() 
		if err != nil {
			respondWithError(w, 401,"cinquième erreur")
			return
		}  
		
			 id, err := strconv.Atoi(Id1)
			 if err != nil {
			
				w.WriteHeader(500)
				return
			 }
		  user:=database.User{
			Id: id,
			Email : param.Email,
			Password : param.Password,
		 }

		 myResponseId, err:=cfg.db.UpdateUser(user)
		 if err != nil {
			
			w.WriteHeader(401)
			return
	  }

	  respondWithJSONUserUpdate(w, 200, myResponseId)
           
	}


	func (cfg *dbConfig) handlerPostRefresh(w http.ResponseWriter, r *http.Request) {
		
		claimsStruct := jwt.RegisteredClaims{}
		headers:=r.Header.Get("Authorization")
		if headers == "" {
			respondWithError(w, 500,"deuxième erreur")
			return
		}
		tokenString:=strings.Split(headers, " ")
		if len(tokenString) < 2 || tokenString[0] != "Bearer" {
			respondWithError(w, 500,"troisieme erreur")
			return
		}

		_, err := jwt.ParseWithClaims(tokenString[1], &claimsStruct, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.config.jwtRefreshSecret), nil})
		if err != nil {
			respondWithError(w, 401,"Quatrième erreur")
			return
		} 
         
		vall, errRevoque :=cfg.RefreshTuu(tokenString[1])

		if errRevoque != nil {
			respondWithError(w, 401,"Quatrième erreur")
			return
		} 

		respondWithJSONUserRefresh(w, 200, vall)
	}



	func (cfg *dbConfig)handlerPostRevoque(w http.ResponseWriter, r *http.Request){
	   
		//claimsStruct := jwt.RegisteredClaims{}
		headers:=r.Header.Get("Authorization")
		if headers == "" {
			respondWithError(w, 500,"deuxième erreur")
			return
		}
		tokenString:=strings.Split(headers, " ")
		if len(tokenString) < 2 || tokenString[0] != "Bearer" {
			respondWithError(w, 500,"troisieme erreur")
			return
		}

		mmm, errRevoque :=cfg.RefreshToken(tokenString[1])

		if errRevoque!=nil {
			respondWithError(w, 401,"Something went wrong")
			return
		}
		respondWithJSONUserRevoke(w, 200, mmm)

	}


	func (cfg *dbConfig) handlerDelete(w http.ResponseWriter, r *http.Request){
		
		idString:=chi.URLParam(r, "chirpId") 
		id, err := strconv.Atoi(idString)
		if err != nil {
		 
		  w.WriteHeader(500)
		  return
	  }
	  id,stringErr:=cfg.handlerGetIdUserForHeaders(r)
	  if id==0 {
		  respondWithError(w, 403,stringErr)
		  return
	  }

		result,err:=cfg.db.DeleteChirp(id)
		if err!=nil {
		   respondWithError(w, 403,"Something went wrong ")
		   return
		}
		dat, err := json.Marshal(result) //(*params)[index]
			if err !=nil {
					
					w.WriteHeader(500)
					return
			}
		
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(dat)  

		}



	func (cfg *dbConfig) handlerPostWebhooks(w http.ResponseWriter, r *http.Request){
        
		headers:=r.Header.Get("Authorization")
		if headers == "" {
			respondWithError(w, 401,"deuxième erreur")
			return
		}
		keyString:=strings.Split(headers, " ")
		if len(keyString) < 2 || keyString[0] != "ApiKey" {
			respondWithError(w, 401,"troisieme erreur")
			return
		}
		
		if keyString[1]!=cfg.config.apiKey {
			respondWithError(w, 401,"deuxième erreur")
			return
		}
		decoder := json.NewDecoder(r.Body)
		param := database.Webhooks{}
		err := decoder.Decode(&param)
		
		if err != nil {
            log.Printf("blublu1")
			respondWithError(w, 404,"Something went wrong")
			return
		}
		log.Println("nono",param)
		result, err:=cfg.db.UpdateUserWebhooks(param)

		if err != nil {
			log.Printf("blublu2")
			respondWithError(w, 404,"Something went wrong")
			return
		}
         a:=database.User{}
		if result ==a {
			w.WriteHeader(200)
			return
		}

		 respondWithJSONUser(w,200, a)

	}
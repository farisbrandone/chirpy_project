package database

import (
	
)

type UserId struct {
	UserID int `json:"user_id"`
}

 type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsChirpyRed bool  `json:"is_chirpy_red"` 
} 

type UserEmailPassword struct {
	Email string `json:"email"`
	Password string `json:"password"`
	ExpiresInSeconds int `json:"expires_in_seconds"`
}

type Webhooks struct {
	Data UserId  `json:"data"`
	Event string `json:"event"`
	
}

/* func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:    user.ID,
		Email: user.Email,
	})
} */

func (db *DB) CreateUser(body UserEmailPassword) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	id := len(dbStructure.Users) + 1
	 user:= User{
		Id:   id,
		Email: body.Email,
		Password: body.Password,
		IsChirpyRed : false,
	}
	dbStructure.Users[id] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) LoginUser(body UserEmailPassword) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user,err1:=EmailPasswordVerification(body,dbStructure.Users)
	
	return user,err1
}

func (db *DB) UpdateUser(user User) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	users,err1:=db.IdUpdate(user,dbStructure)
	
	return users,err1
}

func (db *DB) UpdateUserWebhooks(userWebhoooks Webhooks) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
  a:=User{}
	users,err1:=db.IdUpdateWebhooks(userWebhoooks,dbStructure)
	if err1 != nil {
		return User{}, err1
	}
	if users==a {
		return a, nil
	}
	
	return users,nil
}
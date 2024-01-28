package database

import (
	"os"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	
)

func (db *DB)IdUpdate(user User, dbStructures DBStructure) (User, error){
	
	for key, _:=range dbStructures.Users{
		
		if key == user.Id {
		
		 bcriptPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),14)
		if err!=nil {
			return User{}, err
		}
		userUpdate:=User{
			Id: user.Id,
	        Email : user.Email,
	        Password : string(bcriptPassword),
			IsChirpyRed:dbStructures.Users[key].IsChirpyRed,
		}
		delete(dbStructures.Users, key)
		 dbStructures.Users[key]=userUpdate
          err4 := os.Truncate(db.path, 0)
		if err4 != nil {
			return User{}, err4
		}
		err2 :=db.writeDB( dbStructures)
		if err2 != nil {
			return User{}, err2
		}
		return userUpdate, nil
		}
  }
  return User{}, errors.New("your token does'nt verified")

}


func (db *DB)IdUpdateWebhooks(userWebhoooks Webhooks, dbStructures DBStructure) (User, error){
	
	for key, _:=range dbStructures.Users{
		log.Println("je suis ", userWebhoooks.Data.UserID)
		
		if key == userWebhoooks.Data.UserID {
			log.Println("dedans")
		if userWebhoooks.Event != "user.upgraded" {
			return User{}, nil
		}
		
		userUpdate:=User{
			Id: dbStructures.Users[key].Id,
	        Email : dbStructures.Users[key].Email,
	        Password : dbStructures.Users[key].Password,
			IsChirpyRed:true,
		}
		delete(dbStructures.Users, key)
		 dbStructures.Users[key]=userUpdate
          err4 := os.Truncate(db.path, 0)
		if err4 != nil {
			return User{}, err4
		}
		err2 :=db.writeDB( dbStructures)
		if err2 != nil {
			return User{}, err2
		}
		return userUpdate, nil
		}
  }
  return User{}, errors.New("your token does'nt verified")

}
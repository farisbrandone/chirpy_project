package database

import(
     "errors"
	 "golang.org/x/crypto/bcrypt"
	
)

func EmailPasswordVerification(emailForVerification UserEmailPassword ,users map[int]User) (User, error){
	for key, value:=range users{
		
		
		if value.Email == emailForVerification.Email {
		  passref:=[]byte(value.Password)
		  passwordForCompare:=[]byte(emailForVerification.Password)
		  err:=bcrypt.CompareHashAndPassword(passref,passwordForCompare)
		  
		  if err!=nil {
			  
			  return User{}, errors.New("password is not correct")
		  }
		  return users[key], nil

		}
  }
  return User{}, errors.New("this email doesn't exist")
}
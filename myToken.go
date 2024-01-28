package main

import (
	"github.com/farisbrandone/chirpy_project/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
	
)
type ResToken struct {
		Id int `json:"id"`
		Email string `json:"email"` 
		IsChirpyRed bool  `json:"is_chirpy_red"` 
		Token string  `json:"token"`
		RefreshToken string `json:"refresh_token"`
}

func (cfg *dbConfig)GetToken(param database.UserEmailPassword, PasswordEmailCompare database.User) (ResToken, error) {
	defaultExpiration := 60 * 60 * 24
	if param.ExpiresInSeconds == 0 {
		param.ExpiresInSeconds = defaultExpiration
	} else if param.ExpiresInSeconds > defaultExpiration {
		param.ExpiresInSeconds = defaultExpiration
	}
    
	//token, err := auth.MakeJWT(PasswordEmailCompare.ID, cfg.config.jwtSecret, time.Duration(params.ExpiresInSeconds)*time.Second)
	claims := jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(param.ExpiresInSeconds)*time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1)*time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy-access",
			Subject:   fmt.Sprintf("%d", PasswordEmailCompare.Id),
		}

		claimsRefresh := jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(param.ExpiresInSeconds)*time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(60*24)*time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy-refresh",
			Subject:   fmt.Sprintf("%d", PasswordEmailCompare.Id),
		}
	
	
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	myToken, err:=accessToken.SignedString([]byte(cfg.config.jwtSecret))

	if err!=nil {
		return ResToken{}, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	myTokenRefresh, erra:=refreshToken.SignedString([]byte(cfg.config.jwtRefreshSecret))
	if erra!=nil {
		return ResToken{}, erra
	}
	responseToken:=ResToken{
		Id:PasswordEmailCompare.Id,
		Email : PasswordEmailCompare.Email,
		IsChirpyRed:PasswordEmailCompare.IsChirpyRed,
		Token : myToken,
		RefreshToken: myTokenRefresh,
	}
	return responseToken, nil
}



func (cfg *dbConfig)RefreshToken(param string) (string, error) {
	
	claimsStruct := jwt.RegisteredClaims{}
	
	token, err := jwt.ParseWithClaims(param, &claimsStruct, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.config.jwtRefreshSecret), nil})
	if err != nil {
		return "", err
	} 
	Id1, err := token.Claims.GetSubject() 
	
	if err != nil {
		return "", err
	}  

		claimsRefresh := jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(param.ExpiresInSeconds)*time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(60*24)*time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy-refresh",
			Subject: Id1,
		}
	
	refreshToken1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	mymy, erra:= refreshToken1.SignedString([]byte(cfg.config.jwtRefreshSecret))
	if erra!=nil {
		return "", erra
	}
	
	return mymy, nil
}


func (cfg *dbConfig)RefreshTuu(param string) (string, error) {
	
	claimsStruct := jwt.RegisteredClaims{}
	
	token, err := jwt.ParseWithClaims(param, &claimsStruct, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.config.jwtRefreshSecret), nil})
	if err != nil {
		return "", err
	} 
	Id1, err := token.Claims.GetSubject() 
	if err != nil {
		return "", err
	}  

		claims := jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(param.ExpiresInSeconds)*time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(60*24)*time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer: "chirpy-refresh",
			Subject: Id1,
		}
	
	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mymy, erra:= token1.SignedString([]byte(cfg.config.jwtSecret))
	if erra!=nil {
		return "", erra
	}
	
	return mymy, nil
}
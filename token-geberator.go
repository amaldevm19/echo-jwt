package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateTokenPair() (map[string]string, error) {

	//create Token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = "Jon Doe"
	claims["admin"] = true
	claims["sub"] = 1
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, rterr := refreshToken.SignedString([]byte("secret"))
	if rterr != nil {
		return nil, rterr
	}
	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil

}

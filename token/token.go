package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"jwt-auth.com/config"
	"jwt-auth.com/types"
	"time"
)

func Generatetokens(email *string, username *string, uid string) (signedtoken string, signedrefreshtoken string, err error) {

	claims := &types.SignedDetails{
		Email:    email,
		Username: username,
		User_id:  uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &types.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Getconfig().SecretKey))
	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(config.Getconfig().SecretKey))

	return token, refreshtoken, nil
}

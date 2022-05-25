package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Config struct {
	Port      string
	MongoUrl  string
	SecretKey string
}

type LoginUser struct {
	Email    string
	Password string
}

type SignedDetails struct {
	Email    *string
	Username *string
	User_id  string
	jwt.StandardClaims
}

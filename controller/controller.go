package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-auth.com/Password"
	"jwt-auth.com/models"
	"jwt-auth.com/types"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-auth.com/database"
	"net/http"
	"time"
)

var JwtCollection *mongo.Collection = database.OpenCollection(database.Client, "Userdata")

func Register(c echo.Context) error {
	//handle request
	var user = new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "user Binding failed")
	}
	//set create time
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	//hashpassword

	hashpass := Password.Encrypt(user.Password)
	user.Password = hashpass

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	result, err := JwtCollection.InsertOne(ctx, user)
	defer cancel()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func Login(c echo.Context) error {
	//declare
	var user = new(types.LoginUser)
	var userdata = new(models.User)
	//get login details to user
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "binding login failed")
	}
	//get the password from db

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	err := JwtCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(userdata)
	defer cancel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//check password before giving token
	check := Password.Verifypassword(user.Password, userdata.Password)
	if !check {
		return c.JSON(http.StatusInternalServerError, "failed to verify password")
	}

	return c.JSON(http.StatusOK, "user log in")
}

func Details(c echo.Context) error {
	return c.JSON(http.StatusOK, "view details")
}

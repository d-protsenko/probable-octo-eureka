package repositories

import (
	"awesomeProject/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userProfileCollection *mongo.Collection

func getUserBson(user models.User) bson.M {
	return bson.M{
		"_id":      user.Id,
		"email":    user.Email,
		"password": user.Pwd,
	}
}

func InsertUser(user models.User) error {
	userProfileCollection := MainDB.Collection("user_profile")
	if res, err := userProfileCollection.InsertOne(
		Ctx, getUserBson(user),
	); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.InsertedID)
	}
	return nil
}

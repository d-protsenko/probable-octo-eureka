package repositories

import (
	"awesomeProject/cfg"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Client *mongo.Client
var MainDB *mongo.Database
var Ctx context.Context

func ConnectToDb() {
	credentials := options.Credential{
		Username: cfg.MainDBConfig.User,
		Password: cfg.MainDBConfig.Password,
	}
	client, err := mongo.NewClient(
		options.Client().SetAuth(credentials).ApplyURI("mongodb://localhost:27017"),
	)
	Client = client
	if err != nil {
		fmt.Println("Unable to create new mongo client. Err: ", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Ctx = ctx
	err = Client.Connect(Ctx)
	if err != nil {
		fmt.Println("Unable to connect to the DB. Err: ", err)
	}
	MainDB = Client.Database("awesome_db")
}

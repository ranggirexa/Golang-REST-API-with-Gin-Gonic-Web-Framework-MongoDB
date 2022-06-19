package main

import (
	"context"
	"fmt"
	"log"

	"example.com/sarang-apis/controller"
	"example.com/sarang-apis/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	shoeservice    services.ShoeService
	usercontroller controller.UserController
	shoecontroller controller.ShoeController
	ctx            context.Context
	usercollection *mongo.Collection
	shoecollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection establised")

	usercollection = mongoclient.Database("userdb").Collection("users")
	shoecollection = mongoclient.Database("userdb").Collection("shoes")
	userservice = services.NewUserService(usercollection, ctx)
	shoeservice = services.NewShoeService(shoecollection, ctx)
	usercontroller = controller.New(userservice)
	shoecontroller = controller.NewShoeServicew(shoeservice)
	server = gin.Default()
}

//v1/user/create
func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
	shoecontroller.RegisterShoeRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}

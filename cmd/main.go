package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"simple-blogging-platform/internal"
	"simple-blogging-platform/mongo_client"
	"simple-blogging-platform/routers"
	"simple-blogging-platform/server"
)

func main() {
	ginserver := gin.Default()

	ctx := context.Background()
	log := log.Logger{}

	//connect to mongodb server
	client, err := mongo_client.Connect("mongodb://user:pass@localhost:27017")
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to connect to mongo db"))
	}
	defer func(client *mongo.Client) {
		mongo_client.Close(client)
	}(client)

	server := server.NewServer(ctx, log, client, ginserver)
	blog_client := internal.Blog{
		Ctx:          server.Ctx,
		Mongo_client: server.Mongoclient,
	}
	routers.ApiRouters(ginserver, blog_client)
	ginserver.Run(":8080")
}

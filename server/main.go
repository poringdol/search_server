package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"yandex-food/server/http"
	"yandex-food/utils"
)

const mongoPath = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	signalHandler()

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoPath))
	utils.CheckError(err, "create connection with mongo")
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	switch len(os.Args) {
	case 1:
		startTcp()
	case 2:
		if os.Args[1] == "http" {
			http.StartHttp()
			break
		}
		fallthrough
	default:
		fmt.Println("use \"search_server\" without arguments to start simple tcp-server for search by phone or \"search_server http\" for http-server")
	}
}

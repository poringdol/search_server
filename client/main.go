package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"yandex-food/server/db"
	"yandex-food/utils"
)

const mongoPath = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoPath))
	utils.CheckError(err, "create connection with mongo")
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var phone int64
	if len(os.Args) > 1 {
		switch {
		case os.Args[1] == "--init":
			db.PrepareFiles()
			db.SaveDataToDB(ctx, client)
			db.CreateIndex(ctx, client)
			return
		case os.Args[1] == "--enrich":
			city := ""
			if len(os.Args) > 2 {
				city = os.Args[2]
			}
			db.EnrichAddress(ctx, client, city)
		default:
			intPhone, err := strconv.Atoi(os.Args[1])
			utils.CheckError(err, "get it from sting")
			phone = int64(intPhone)
		}
	} else {
		fmt.Print("Enter phone number: ")
		fmt.Scanf("%d", &phone)
		fmt.Println()
	}

	byPhone, err := db.GetByPhone(ctx, client, phone)
	utils.CheckError(err, "get records by phone")

	fmt.Printf("Список адресов для телефона %d:\n", phone)
	for _, rec := range byPhone {
		fmt.Println("\t", rec.AddressCity, rec.AddressStreet, rec.AddressHouse, rec.AddressFloor, rec.AddressOffice)
	}

	for _, rec := range byPhone {
		byAddress, err := db.GetByAddress(ctx, client, rec)
		utils.CheckError(err, "get record by address")
		fmt.Println("=")
		for _, recByAddr := range byAddress {
			utils.PrintRecord(os.Stdout, recByAddr)
		}
		fmt.Println("================================================")
	}
}
